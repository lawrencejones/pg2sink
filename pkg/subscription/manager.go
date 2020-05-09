package subscription

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lawrencejones/pg2sink/pkg/util"

	kitlog "github.com/go-kit/kit/log"
	"github.com/jackc/pgtype"
	"github.com/pkg/errors"
)

type ManagerOptions struct {
	Name         string        // name of the publication in Postgres
	Schemas      []string      // list of schemas to watch
	Excludes     []string      // optional blacklist
	Includes     []string      // optional whitelist, combined with blacklist
	PollInterval time.Duration // interval to scan Postgres for new matching tables
}

// Manager supervises a subscription, adding and removing tables into the publication
// according to the white/blacklist.
type Manager struct {
	logger kitlog.Logger
	conn   querier
	opts   ManagerOptions
}

func NewManager(logger kitlog.Logger, conn querier, opts ManagerOptions) *Manager {
	return &Manager{
		logger: logger,
		conn:   conn,
		opts:   opts,
	}
}

// Manage begins syncing tables into publication using the rules configured on the manager
// options. It will run until the context expires.
func (m *Manager) Manage(ctx context.Context) (err error) {
	logger := kitlog.With(m.logger, "publication_name", m.opts.Name)
	publication := &Publication{Name: m.opts.Name}

	for {
		logger.Log("event", "sync_published_tables")
		watched, err := m.getWatchedTables(ctx)
		if err != nil {
			return errors.Wrap(err, "failed to discover watched tables")
		}

		published, err := publication.GetTables(ctx, m.conn)
		if err != nil {
			return errors.Wrap(err, "failed to query published tables")
		}

		watchedNotPublished := util.Diff(watched, published)
		publishedNotWatched := util.Diff(published, watched)

		if (len(watchedNotPublished) > 0) || (len(publishedNotWatched) > 0) {
			logger.Log("event", "alter_publication",
				"adding", strings.Join(watchedNotPublished, ","),
				"removing", strings.Join(publishedNotWatched, ","))

			if err := publication.SetTables(ctx, m.conn, watched...); err != nil {
				return fmt.Errorf("failed to alter publication: %w", err)
			}
		}

		select {
		case <-ctx.Done():
			logger.Log("event", "finish", "msg", "context expired, finishing sync")
			return nil
		case <-time.After(m.opts.PollInterval):
			// continue
		}
	}
}

// getWatchedTables scans the database for tables that match our watch conditions
func (p *Manager) getWatchedTables(ctx context.Context) ([]string, error) {
	query := `
	select array_agg(table_schema || '.' || table_name)
	from information_schema.tables
	where table_schema = any($1::text[])
	and table_type = 'BASE TABLE';
	`

	tablesReceiver := pgtype.TextArray{}
	var tables []string
	if err := p.conn.QueryRow(ctx, query, p.opts.Schemas).Scan(&tablesReceiver); err != nil {
		return nil, err
	}

	if err := tablesReceiver.AssignTo(&tables); err != nil {
		return nil, err
	}

	watchedTables := []string{}
forEachRow:
	for _, table := range tables {
		if util.Includes(p.opts.Excludes, table) {
			continue forEachRow
		}

		// If we have an includes list, we only include our table if it appears in that list
		isIncluded := len(p.opts.Includes) == 0
		isIncluded = isIncluded || util.Includes(p.opts.Includes, table)

		if !isIncluded {
			continue forEachRow
		}

		watchedTables = append(watchedTables, table)
	}

	return watchedTables, nil
}

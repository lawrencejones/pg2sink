package imports

import (
	"context"
	"time"

	kitlog "github.com/go-kit/kit/log"
	"github.com/jackc/pgx"
	"github.com/lawrencejones/pg2sink/pkg/publication"
	"github.com/lawrencejones/pg2sink/pkg/util"
	"github.com/pkg/errors"
)

type ManagerOptions struct {
	PublicationID string        // identifier of the current publication
	PollInterval  time.Duration // interval to poll for new import jobs
}

func NewManager(logger kitlog.Logger, pool *pgx.ConnPool, opts ManagerOptions) *Manager {
	return &Manager{
		logger: logger,
		pool:   pool,
		opts:   opts,
	}
}

// Manager ensures we create import jobs for all tables that have not yet been imported.
type Manager struct {
	logger kitlog.Logger
	pool   *pgx.ConnPool
	opts   ManagerOptions
}

// Work starts WorkerCount workers to process jobs in the import jobs table. It returns a
// channel of Committed structs, generated from the on-going imports. It will track
// progress in the import jobs table.

// Sync watches the publication to identify tables which have not yet been imported, and
// adds import jobs for those tables.
func (i Manager) Sync(ctx context.Context) error {
	logger := kitlog.With(i.logger, "publication_id", i.opts.PublicationID)

	for {
		logger.Log("event", "sync.poll")
		publishedTables, err := publication.GetPublishedTables(ctx, i.pool, i.opts.PublicationID)
		if err != nil {
			return errors.Wrap(err, "failed to query published tables")
		}

		importedTables, err := JobStore{i.pool}.GetImportedTables(ctx, i.opts.PublicationID)
		if err != nil {
			return errors.Wrap(err, "failed to query imported tables")
		}

		notImportedTables := util.Diff(publishedTables, importedTables)
		for _, table := range notImportedTables {
			logger.Log("event", "import_job.create", "table", table)
			job, err := JobStore{i.pool}.Create(ctx, i.opts.PublicationID, table)
			if err != nil {
				return errors.Wrap(err, "failed to create import job")
			}

			logger.Log("event", "import_job.created", "table", table, "job_id", job.ID)
		}

		select {
		case <-ctx.Done():
			logger.Log("event", "sync.finish", "msg", "context expired, finishing sync")
			return nil
		case <-time.After(i.opts.PollInterval):
			// continue
		}
	}
}
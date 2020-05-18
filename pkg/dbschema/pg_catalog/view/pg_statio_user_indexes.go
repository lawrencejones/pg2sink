//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/postgres"
)

var PgStatioUserIndexes = newPgStatioUserIndexesTable()

type PgStatioUserIndexesTable struct {
	postgres.Table

	//Columns
	Relid        postgres.ColumnString
	Indexrelid   postgres.ColumnString
	Schemaname   postgres.ColumnString
	Relname      postgres.ColumnString
	Indexrelname postgres.ColumnString
	IdxBlksRead  postgres.ColumnInteger
	IdxBlksHit   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatioUserIndexesTable with assigned alias
func (a *PgStatioUserIndexesTable) AS(alias string) *PgStatioUserIndexesTable {
	aliasTable := newPgStatioUserIndexesTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatioUserIndexesTable() *PgStatioUserIndexesTable {
	var (
		RelidColumn        = postgres.StringColumn("relid")
		IndexrelidColumn   = postgres.StringColumn("indexrelid")
		SchemanameColumn   = postgres.StringColumn("schemaname")
		RelnameColumn      = postgres.StringColumn("relname")
		IndexrelnameColumn = postgres.StringColumn("indexrelname")
		IdxBlksReadColumn  = postgres.IntegerColumn("idx_blks_read")
		IdxBlksHitColumn   = postgres.IntegerColumn("idx_blks_hit")
	)

	return &PgStatioUserIndexesTable{
		Table: postgres.NewTable("pg_catalog", "pg_statio_user_indexes", RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxBlksReadColumn, IdxBlksHitColumn),

		//Columns
		Relid:        RelidColumn,
		Indexrelid:   IndexrelidColumn,
		Schemaname:   SchemanameColumn,
		Relname:      RelnameColumn,
		Indexrelname: IndexrelnameColumn,
		IdxBlksRead:  IdxBlksReadColumn,
		IdxBlksHit:   IdxBlksHitColumn,

		AllColumns:     postgres.ColumnList{RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxBlksReadColumn, IdxBlksHitColumn},
		MutableColumns: postgres.ColumnList{RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxBlksReadColumn, IdxBlksHitColumn},
	}
}

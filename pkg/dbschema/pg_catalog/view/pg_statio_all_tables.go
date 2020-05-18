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

var PgStatioAllTables = newPgStatioAllTablesTable()

type PgStatioAllTablesTable struct {
	postgres.Table

	//Columns
	Relid         postgres.ColumnString
	Schemaname    postgres.ColumnString
	Relname       postgres.ColumnString
	HeapBlksRead  postgres.ColumnInteger
	HeapBlksHit   postgres.ColumnInteger
	IdxBlksRead   postgres.ColumnInteger
	IdxBlksHit    postgres.ColumnInteger
	ToastBlksRead postgres.ColumnInteger
	ToastBlksHit  postgres.ColumnInteger
	TidxBlksRead  postgres.ColumnInteger
	TidxBlksHit   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatioAllTablesTable with assigned alias
func (a *PgStatioAllTablesTable) AS(alias string) *PgStatioAllTablesTable {
	aliasTable := newPgStatioAllTablesTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatioAllTablesTable() *PgStatioAllTablesTable {
	var (
		RelidColumn         = postgres.StringColumn("relid")
		SchemanameColumn    = postgres.StringColumn("schemaname")
		RelnameColumn       = postgres.StringColumn("relname")
		HeapBlksReadColumn  = postgres.IntegerColumn("heap_blks_read")
		HeapBlksHitColumn   = postgres.IntegerColumn("heap_blks_hit")
		IdxBlksReadColumn   = postgres.IntegerColumn("idx_blks_read")
		IdxBlksHitColumn    = postgres.IntegerColumn("idx_blks_hit")
		ToastBlksReadColumn = postgres.IntegerColumn("toast_blks_read")
		ToastBlksHitColumn  = postgres.IntegerColumn("toast_blks_hit")
		TidxBlksReadColumn  = postgres.IntegerColumn("tidx_blks_read")
		TidxBlksHitColumn   = postgres.IntegerColumn("tidx_blks_hit")
	)

	return &PgStatioAllTablesTable{
		Table: postgres.NewTable("pg_catalog", "pg_statio_all_tables", RelidColumn, SchemanameColumn, RelnameColumn, HeapBlksReadColumn, HeapBlksHitColumn, IdxBlksReadColumn, IdxBlksHitColumn, ToastBlksReadColumn, ToastBlksHitColumn, TidxBlksReadColumn, TidxBlksHitColumn),

		//Columns
		Relid:         RelidColumn,
		Schemaname:    SchemanameColumn,
		Relname:       RelnameColumn,
		HeapBlksRead:  HeapBlksReadColumn,
		HeapBlksHit:   HeapBlksHitColumn,
		IdxBlksRead:   IdxBlksReadColumn,
		IdxBlksHit:    IdxBlksHitColumn,
		ToastBlksRead: ToastBlksReadColumn,
		ToastBlksHit:  ToastBlksHitColumn,
		TidxBlksRead:  TidxBlksReadColumn,
		TidxBlksHit:   TidxBlksHitColumn,

		AllColumns:     postgres.ColumnList{RelidColumn, SchemanameColumn, RelnameColumn, HeapBlksReadColumn, HeapBlksHitColumn, IdxBlksReadColumn, IdxBlksHitColumn, ToastBlksReadColumn, ToastBlksHitColumn, TidxBlksReadColumn, TidxBlksHitColumn},
		MutableColumns: postgres.ColumnList{RelidColumn, SchemanameColumn, RelnameColumn, HeapBlksReadColumn, HeapBlksHitColumn, IdxBlksReadColumn, IdxBlksHitColumn, ToastBlksReadColumn, ToastBlksHitColumn, TidxBlksReadColumn, TidxBlksHitColumn},
	}
}

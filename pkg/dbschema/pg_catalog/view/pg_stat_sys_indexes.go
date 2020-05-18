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

var PgStatSysIndexes = newPgStatSysIndexesTable()

type PgStatSysIndexesTable struct {
	postgres.Table

	//Columns
	Relid        postgres.ColumnString
	Indexrelid   postgres.ColumnString
	Schemaname   postgres.ColumnString
	Relname      postgres.ColumnString
	Indexrelname postgres.ColumnString
	IdxScan      postgres.ColumnInteger
	IdxTupRead   postgres.ColumnInteger
	IdxTupFetch  postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatSysIndexesTable with assigned alias
func (a *PgStatSysIndexesTable) AS(alias string) *PgStatSysIndexesTable {
	aliasTable := newPgStatSysIndexesTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatSysIndexesTable() *PgStatSysIndexesTable {
	var (
		RelidColumn        = postgres.StringColumn("relid")
		IndexrelidColumn   = postgres.StringColumn("indexrelid")
		SchemanameColumn   = postgres.StringColumn("schemaname")
		RelnameColumn      = postgres.StringColumn("relname")
		IndexrelnameColumn = postgres.StringColumn("indexrelname")
		IdxScanColumn      = postgres.IntegerColumn("idx_scan")
		IdxTupReadColumn   = postgres.IntegerColumn("idx_tup_read")
		IdxTupFetchColumn  = postgres.IntegerColumn("idx_tup_fetch")
	)

	return &PgStatSysIndexesTable{
		Table: postgres.NewTable("pg_catalog", "pg_stat_sys_indexes", RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxScanColumn, IdxTupReadColumn, IdxTupFetchColumn),

		//Columns
		Relid:        RelidColumn,
		Indexrelid:   IndexrelidColumn,
		Schemaname:   SchemanameColumn,
		Relname:      RelnameColumn,
		Indexrelname: IndexrelnameColumn,
		IdxScan:      IdxScanColumn,
		IdxTupRead:   IdxTupReadColumn,
		IdxTupFetch:  IdxTupFetchColumn,

		AllColumns:     postgres.ColumnList{RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxScanColumn, IdxTupReadColumn, IdxTupFetchColumn},
		MutableColumns: postgres.ColumnList{RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxScanColumn, IdxTupReadColumn, IdxTupFetchColumn},
	}
}

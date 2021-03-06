//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PgStatUserIndexes = newPgStatUserIndexesTable()

type pgStatUserIndexesTable struct {
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

type PgStatUserIndexesTable struct {
	pgStatUserIndexesTable

	EXCLUDED pgStatUserIndexesTable
}

// AS creates new PgStatUserIndexesTable with assigned alias
func (a *PgStatUserIndexesTable) AS(alias string) *PgStatUserIndexesTable {
	aliasTable := newPgStatUserIndexesTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgStatUserIndexesTable() *PgStatUserIndexesTable {
	return &PgStatUserIndexesTable{
		pgStatUserIndexesTable: newPgStatUserIndexesTableImpl("pg_catalog", "pg_stat_user_indexes"),
		EXCLUDED:               newPgStatUserIndexesTableImpl("", "excluded"),
	}
}

func newPgStatUserIndexesTableImpl(schemaName, tableName string) pgStatUserIndexesTable {
	var (
		RelidColumn        = postgres.StringColumn("relid")
		IndexrelidColumn   = postgres.StringColumn("indexrelid")
		SchemanameColumn   = postgres.StringColumn("schemaname")
		RelnameColumn      = postgres.StringColumn("relname")
		IndexrelnameColumn = postgres.StringColumn("indexrelname")
		IdxScanColumn      = postgres.IntegerColumn("idx_scan")
		IdxTupReadColumn   = postgres.IntegerColumn("idx_tup_read")
		IdxTupFetchColumn  = postgres.IntegerColumn("idx_tup_fetch")
		allColumns         = postgres.ColumnList{RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxScanColumn, IdxTupReadColumn, IdxTupFetchColumn}
		mutableColumns     = postgres.ColumnList{RelidColumn, IndexrelidColumn, SchemanameColumn, RelnameColumn, IndexrelnameColumn, IdxScanColumn, IdxTupReadColumn, IdxTupFetchColumn}
	)

	return pgStatUserIndexesTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Relid:        RelidColumn,
		Indexrelid:   IndexrelidColumn,
		Schemaname:   SchemanameColumn,
		Relname:      RelnameColumn,
		Indexrelname: IndexrelnameColumn,
		IdxScan:      IdxScanColumn,
		IdxTupRead:   IdxTupReadColumn,
		IdxTupFetch:  IdxTupFetchColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

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

var PgIndexes = newPgIndexesTable()

type PgIndexesTable struct {
	postgres.Table

	//Columns
	Schemaname postgres.ColumnString
	Tablename  postgres.ColumnString
	Indexname  postgres.ColumnString
	Tablespace postgres.ColumnString
	Indexdef   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgIndexesTable with assigned alias
func (a *PgIndexesTable) AS(alias string) *PgIndexesTable {
	aliasTable := newPgIndexesTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgIndexesTable() *PgIndexesTable {
	var (
		SchemanameColumn = postgres.StringColumn("schemaname")
		TablenameColumn  = postgres.StringColumn("tablename")
		IndexnameColumn  = postgres.StringColumn("indexname")
		TablespaceColumn = postgres.StringColumn("tablespace")
		IndexdefColumn   = postgres.StringColumn("indexdef")
	)

	return &PgIndexesTable{
		Table: postgres.NewTable("pg_catalog", "pg_indexes", SchemanameColumn, TablenameColumn, IndexnameColumn, TablespaceColumn, IndexdefColumn),

		//Columns
		Schemaname: SchemanameColumn,
		Tablename:  TablenameColumn,
		Indexname:  IndexnameColumn,
		Tablespace: TablespaceColumn,
		Indexdef:   IndexdefColumn,

		AllColumns:     postgres.ColumnList{SchemanameColumn, TablenameColumn, IndexnameColumn, TablespaceColumn, IndexdefColumn},
		MutableColumns: postgres.ColumnList{SchemanameColumn, TablenameColumn, IndexnameColumn, TablespaceColumn, IndexdefColumn},
	}
}

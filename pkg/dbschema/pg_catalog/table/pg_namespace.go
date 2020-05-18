//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/postgres"
)

var PgNamespace = newPgNamespaceTable()

type PgNamespaceTable struct {
	postgres.Table

	//Columns
	Oid      postgres.ColumnString
	Nspname  postgres.ColumnString
	Nspowner postgres.ColumnString
	Nspacl   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgNamespaceTable with assigned alias
func (a *PgNamespaceTable) AS(alias string) *PgNamespaceTable {
	aliasTable := newPgNamespaceTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgNamespaceTable() *PgNamespaceTable {
	var (
		OidColumn      = postgres.StringColumn("oid")
		NspnameColumn  = postgres.StringColumn("nspname")
		NspownerColumn = postgres.StringColumn("nspowner")
		NspaclColumn   = postgres.StringColumn("nspacl")
	)

	return &PgNamespaceTable{
		Table: postgres.NewTable("pg_catalog", "pg_namespace", OidColumn, NspnameColumn, NspownerColumn, NspaclColumn),

		//Columns
		Oid:      OidColumn,
		Nspname:  NspnameColumn,
		Nspowner: NspownerColumn,
		Nspacl:   NspaclColumn,

		AllColumns:     postgres.ColumnList{OidColumn, NspnameColumn, NspownerColumn, NspaclColumn},
		MutableColumns: postgres.ColumnList{OidColumn, NspnameColumn, NspownerColumn, NspaclColumn},
	}
}

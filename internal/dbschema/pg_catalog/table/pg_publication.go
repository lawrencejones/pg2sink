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

var PgPublication = newPgPublicationTable()

type PgPublicationTable struct {
	postgres.Table

	//Columns
	Oid          postgres.ColumnString
	Pubname      postgres.ColumnString
	Pubowner     postgres.ColumnString
	Puballtables postgres.ColumnBool
	Pubinsert    postgres.ColumnBool
	Pubupdate    postgres.ColumnBool
	Pubdelete    postgres.ColumnBool
	Pubtruncate  postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgPublicationTable with assigned alias
func (a *PgPublicationTable) AS(alias string) *PgPublicationTable {
	aliasTable := newPgPublicationTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgPublicationTable() *PgPublicationTable {
	var (
		OidColumn          = postgres.StringColumn("oid")
		PubnameColumn      = postgres.StringColumn("pubname")
		PubownerColumn     = postgres.StringColumn("pubowner")
		PuballtablesColumn = postgres.BoolColumn("puballtables")
		PubinsertColumn    = postgres.BoolColumn("pubinsert")
		PubupdateColumn    = postgres.BoolColumn("pubupdate")
		PubdeleteColumn    = postgres.BoolColumn("pubdelete")
		PubtruncateColumn  = postgres.BoolColumn("pubtruncate")
	)

	return &PgPublicationTable{
		Table: postgres.NewTable("pg_catalog", "pg_publication", OidColumn, PubnameColumn, PubownerColumn, PuballtablesColumn, PubinsertColumn, PubupdateColumn, PubdeleteColumn, PubtruncateColumn),

		//Columns
		Oid:          OidColumn,
		Pubname:      PubnameColumn,
		Pubowner:     PubownerColumn,
		Puballtables: PuballtablesColumn,
		Pubinsert:    PubinsertColumn,
		Pubupdate:    PubupdateColumn,
		Pubdelete:    PubdeleteColumn,
		Pubtruncate:  PubtruncateColumn,

		AllColumns:     postgres.ColumnList{OidColumn, PubnameColumn, PubownerColumn, PuballtablesColumn, PubinsertColumn, PubupdateColumn, PubdeleteColumn, PubtruncateColumn},
		MutableColumns: postgres.ColumnList{OidColumn, PubnameColumn, PubownerColumn, PuballtablesColumn, PubinsertColumn, PubupdateColumn, PubdeleteColumn, PubtruncateColumn},
	}
}
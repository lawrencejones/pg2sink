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

var PgAvailableExtensions = newPgAvailableExtensionsTable()

type PgAvailableExtensionsTable struct {
	postgres.Table

	//Columns
	Name             postgres.ColumnString
	DefaultVersion   postgres.ColumnString
	InstalledVersion postgres.ColumnString
	Comment          postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgAvailableExtensionsTable with assigned alias
func (a *PgAvailableExtensionsTable) AS(alias string) *PgAvailableExtensionsTable {
	aliasTable := newPgAvailableExtensionsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgAvailableExtensionsTable() *PgAvailableExtensionsTable {
	var (
		NameColumn             = postgres.StringColumn("name")
		DefaultVersionColumn   = postgres.StringColumn("default_version")
		InstalledVersionColumn = postgres.StringColumn("installed_version")
		CommentColumn          = postgres.StringColumn("comment")
	)

	return &PgAvailableExtensionsTable{
		Table: postgres.NewTable("pg_catalog", "pg_available_extensions", NameColumn, DefaultVersionColumn, InstalledVersionColumn, CommentColumn),

		//Columns
		Name:             NameColumn,
		DefaultVersion:   DefaultVersionColumn,
		InstalledVersion: InstalledVersionColumn,
		Comment:          CommentColumn,

		AllColumns:     postgres.ColumnList{NameColumn, DefaultVersionColumn, InstalledVersionColumn, CommentColumn},
		MutableColumns: postgres.ColumnList{NameColumn, DefaultVersionColumn, InstalledVersionColumn, CommentColumn},
	}
}

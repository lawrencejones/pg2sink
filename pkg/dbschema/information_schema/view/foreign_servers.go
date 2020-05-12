//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 09:15:06 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/postgres"
)

var ForeignServers = newForeignServersTable()

type ForeignServersTable struct {
	postgres.Table

	//Columns
	ForeignServerCatalog      postgres.ColumnString
	ForeignServerName         postgres.ColumnString
	ForeignDataWrapperCatalog postgres.ColumnString
	ForeignDataWrapperName    postgres.ColumnString
	ForeignServerType         postgres.ColumnString
	ForeignServerVersion      postgres.ColumnString
	AuthorizationIdentifier   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new ForeignServersTable with assigned alias
func (a *ForeignServersTable) AS(alias string) *ForeignServersTable {
	aliasTable := newForeignServersTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newForeignServersTable() *ForeignServersTable {
	var (
		ForeignServerCatalogColumn      = postgres.StringColumn("foreign_server_catalog")
		ForeignServerNameColumn         = postgres.StringColumn("foreign_server_name")
		ForeignDataWrapperCatalogColumn = postgres.StringColumn("foreign_data_wrapper_catalog")
		ForeignDataWrapperNameColumn    = postgres.StringColumn("foreign_data_wrapper_name")
		ForeignServerTypeColumn         = postgres.StringColumn("foreign_server_type")
		ForeignServerVersionColumn      = postgres.StringColumn("foreign_server_version")
		AuthorizationIdentifierColumn   = postgres.StringColumn("authorization_identifier")
	)

	return &ForeignServersTable{
		Table: postgres.NewTable("information_schema", "foreign_servers", ForeignServerCatalogColumn, ForeignServerNameColumn, ForeignDataWrapperCatalogColumn, ForeignDataWrapperNameColumn, ForeignServerTypeColumn, ForeignServerVersionColumn, AuthorizationIdentifierColumn),

		//Columns
		ForeignServerCatalog:      ForeignServerCatalogColumn,
		ForeignServerName:         ForeignServerNameColumn,
		ForeignDataWrapperCatalog: ForeignDataWrapperCatalogColumn,
		ForeignDataWrapperName:    ForeignDataWrapperNameColumn,
		ForeignServerType:         ForeignServerTypeColumn,
		ForeignServerVersion:      ForeignServerVersionColumn,
		AuthorizationIdentifier:   AuthorizationIdentifierColumn,

		AllColumns:     postgres.ColumnList{ForeignServerCatalogColumn, ForeignServerNameColumn, ForeignDataWrapperCatalogColumn, ForeignDataWrapperNameColumn, ForeignServerTypeColumn, ForeignServerVersionColumn, AuthorizationIdentifierColumn},
		MutableColumns: postgres.ColumnList{ForeignServerCatalogColumn, ForeignServerNameColumn, ForeignDataWrapperCatalogColumn, ForeignDataWrapperNameColumn, ForeignServerTypeColumn, ForeignServerVersionColumn, AuthorizationIdentifierColumn},
	}
}

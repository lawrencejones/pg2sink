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

var UserMappingOptions = newUserMappingOptionsTable()

type UserMappingOptionsTable struct {
	postgres.Table

	//Columns
	AuthorizationIdentifier postgres.ColumnString
	ForeignServerCatalog    postgres.ColumnString
	ForeignServerName       postgres.ColumnString
	OptionName              postgres.ColumnString
	OptionValue             postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new UserMappingOptionsTable with assigned alias
func (a *UserMappingOptionsTable) AS(alias string) *UserMappingOptionsTable {
	aliasTable := newUserMappingOptionsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newUserMappingOptionsTable() *UserMappingOptionsTable {
	var (
		AuthorizationIdentifierColumn = postgres.StringColumn("authorization_identifier")
		ForeignServerCatalogColumn    = postgres.StringColumn("foreign_server_catalog")
		ForeignServerNameColumn       = postgres.StringColumn("foreign_server_name")
		OptionNameColumn              = postgres.StringColumn("option_name")
		OptionValueColumn             = postgres.StringColumn("option_value")
	)

	return &UserMappingOptionsTable{
		Table: postgres.NewTable("information_schema", "user_mapping_options", AuthorizationIdentifierColumn, ForeignServerCatalogColumn, ForeignServerNameColumn, OptionNameColumn, OptionValueColumn),

		//Columns
		AuthorizationIdentifier: AuthorizationIdentifierColumn,
		ForeignServerCatalog:    ForeignServerCatalogColumn,
		ForeignServerName:       ForeignServerNameColumn,
		OptionName:              OptionNameColumn,
		OptionValue:             OptionValueColumn,

		AllColumns:     postgres.ColumnList{AuthorizationIdentifierColumn, ForeignServerCatalogColumn, ForeignServerNameColumn, OptionNameColumn, OptionValueColumn},
		MutableColumns: postgres.ColumnList{AuthorizationIdentifierColumn, ForeignServerCatalogColumn, ForeignServerNameColumn, OptionNameColumn, OptionValueColumn},
	}
}
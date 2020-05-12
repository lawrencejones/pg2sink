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

var ForeignDataWrapperOptions = newForeignDataWrapperOptionsTable()

type ForeignDataWrapperOptionsTable struct {
	postgres.Table

	//Columns
	ForeignDataWrapperCatalog postgres.ColumnString
	ForeignDataWrapperName    postgres.ColumnString
	OptionName                postgres.ColumnString
	OptionValue               postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new ForeignDataWrapperOptionsTable with assigned alias
func (a *ForeignDataWrapperOptionsTable) AS(alias string) *ForeignDataWrapperOptionsTable {
	aliasTable := newForeignDataWrapperOptionsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newForeignDataWrapperOptionsTable() *ForeignDataWrapperOptionsTable {
	var (
		ForeignDataWrapperCatalogColumn = postgres.StringColumn("foreign_data_wrapper_catalog")
		ForeignDataWrapperNameColumn    = postgres.StringColumn("foreign_data_wrapper_name")
		OptionNameColumn                = postgres.StringColumn("option_name")
		OptionValueColumn               = postgres.StringColumn("option_value")
	)

	return &ForeignDataWrapperOptionsTable{
		Table: postgres.NewTable("information_schema", "foreign_data_wrapper_options", ForeignDataWrapperCatalogColumn, ForeignDataWrapperNameColumn, OptionNameColumn, OptionValueColumn),

		//Columns
		ForeignDataWrapperCatalog: ForeignDataWrapperCatalogColumn,
		ForeignDataWrapperName:    ForeignDataWrapperNameColumn,
		OptionName:                OptionNameColumn,
		OptionValue:               OptionValueColumn,

		AllColumns:     postgres.ColumnList{ForeignDataWrapperCatalogColumn, ForeignDataWrapperNameColumn, OptionNameColumn, OptionValueColumn},
		MutableColumns: postgres.ColumnList{ForeignDataWrapperCatalogColumn, ForeignDataWrapperNameColumn, OptionNameColumn, OptionValueColumn},
	}
}

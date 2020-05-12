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

var ColumnOptions = newColumnOptionsTable()

type ColumnOptionsTable struct {
	postgres.Table

	//Columns
	TableCatalog postgres.ColumnString
	TableSchema  postgres.ColumnString
	TableName    postgres.ColumnString
	ColumnName   postgres.ColumnString
	OptionName   postgres.ColumnString
	OptionValue  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new ColumnOptionsTable with assigned alias
func (a *ColumnOptionsTable) AS(alias string) *ColumnOptionsTable {
	aliasTable := newColumnOptionsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newColumnOptionsTable() *ColumnOptionsTable {
	var (
		TableCatalogColumn = postgres.StringColumn("table_catalog")
		TableSchemaColumn  = postgres.StringColumn("table_schema")
		TableNameColumn    = postgres.StringColumn("table_name")
		ColumnNameColumn   = postgres.StringColumn("column_name")
		OptionNameColumn   = postgres.StringColumn("option_name")
		OptionValueColumn  = postgres.StringColumn("option_value")
	)

	return &ColumnOptionsTable{
		Table: postgres.NewTable("information_schema", "column_options", TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, OptionNameColumn, OptionValueColumn),

		//Columns
		TableCatalog: TableCatalogColumn,
		TableSchema:  TableSchemaColumn,
		TableName:    TableNameColumn,
		ColumnName:   ColumnNameColumn,
		OptionName:   OptionNameColumn,
		OptionValue:  OptionValueColumn,

		AllColumns:     postgres.ColumnList{TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, OptionNameColumn, OptionValueColumn},
		MutableColumns: postgres.ColumnList{TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, OptionNameColumn, OptionValueColumn},
	}
}

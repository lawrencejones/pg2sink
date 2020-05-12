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

var ViewColumnUsage = newViewColumnUsageTable()

type ViewColumnUsageTable struct {
	postgres.Table

	//Columns
	ViewCatalog  postgres.ColumnString
	ViewSchema   postgres.ColumnString
	ViewName     postgres.ColumnString
	TableCatalog postgres.ColumnString
	TableSchema  postgres.ColumnString
	TableName    postgres.ColumnString
	ColumnName   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new ViewColumnUsageTable with assigned alias
func (a *ViewColumnUsageTable) AS(alias string) *ViewColumnUsageTable {
	aliasTable := newViewColumnUsageTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newViewColumnUsageTable() *ViewColumnUsageTable {
	var (
		ViewCatalogColumn  = postgres.StringColumn("view_catalog")
		ViewSchemaColumn   = postgres.StringColumn("view_schema")
		ViewNameColumn     = postgres.StringColumn("view_name")
		TableCatalogColumn = postgres.StringColumn("table_catalog")
		TableSchemaColumn  = postgres.StringColumn("table_schema")
		TableNameColumn    = postgres.StringColumn("table_name")
		ColumnNameColumn   = postgres.StringColumn("column_name")
	)

	return &ViewColumnUsageTable{
		Table: postgres.NewTable("information_schema", "view_column_usage", ViewCatalogColumn, ViewSchemaColumn, ViewNameColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn),

		//Columns
		ViewCatalog:  ViewCatalogColumn,
		ViewSchema:   ViewSchemaColumn,
		ViewName:     ViewNameColumn,
		TableCatalog: TableCatalogColumn,
		TableSchema:  TableSchemaColumn,
		TableName:    TableNameColumn,
		ColumnName:   ColumnNameColumn,

		AllColumns:     postgres.ColumnList{ViewCatalogColumn, ViewSchemaColumn, ViewNameColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn},
		MutableColumns: postgres.ColumnList{ViewCatalogColumn, ViewSchemaColumn, ViewNameColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn},
	}
}

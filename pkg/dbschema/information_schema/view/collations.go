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

var Collations = newCollationsTable()

type CollationsTable struct {
	postgres.Table

	//Columns
	CollationCatalog postgres.ColumnString
	CollationSchema  postgres.ColumnString
	CollationName    postgres.ColumnString
	PadAttribute     postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new CollationsTable with assigned alias
func (a *CollationsTable) AS(alias string) *CollationsTable {
	aliasTable := newCollationsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newCollationsTable() *CollationsTable {
	var (
		CollationCatalogColumn = postgres.StringColumn("collation_catalog")
		CollationSchemaColumn  = postgres.StringColumn("collation_schema")
		CollationNameColumn    = postgres.StringColumn("collation_name")
		PadAttributeColumn     = postgres.StringColumn("pad_attribute")
	)

	return &CollationsTable{
		Table: postgres.NewTable("information_schema", "collations", CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, PadAttributeColumn),

		//Columns
		CollationCatalog: CollationCatalogColumn,
		CollationSchema:  CollationSchemaColumn,
		CollationName:    CollationNameColumn,
		PadAttribute:     PadAttributeColumn,

		AllColumns:     postgres.ColumnList{CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, PadAttributeColumn},
		MutableColumns: postgres.ColumnList{CollationCatalogColumn, CollationSchemaColumn, CollationNameColumn, PadAttributeColumn},
	}
}

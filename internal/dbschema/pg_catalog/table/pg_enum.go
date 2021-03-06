//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PgEnum = newPgEnumTable()

type pgEnumTable struct {
	postgres.Table

	//Columns
	Oid           postgres.ColumnString
	Enumtypid     postgres.ColumnString
	Enumsortorder postgres.ColumnFloat
	Enumlabel     postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgEnumTable struct {
	pgEnumTable

	EXCLUDED pgEnumTable
}

// AS creates new PgEnumTable with assigned alias
func (a *PgEnumTable) AS(alias string) *PgEnumTable {
	aliasTable := newPgEnumTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgEnumTable() *PgEnumTable {
	return &PgEnumTable{
		pgEnumTable: newPgEnumTableImpl("pg_catalog", "pg_enum"),
		EXCLUDED:    newPgEnumTableImpl("", "excluded"),
	}
}

func newPgEnumTableImpl(schemaName, tableName string) pgEnumTable {
	var (
		OidColumn           = postgres.StringColumn("oid")
		EnumtypidColumn     = postgres.StringColumn("enumtypid")
		EnumsortorderColumn = postgres.FloatColumn("enumsortorder")
		EnumlabelColumn     = postgres.StringColumn("enumlabel")
		allColumns          = postgres.ColumnList{OidColumn, EnumtypidColumn, EnumsortorderColumn, EnumlabelColumn}
		mutableColumns      = postgres.ColumnList{OidColumn, EnumtypidColumn, EnumsortorderColumn, EnumlabelColumn}
	)

	return pgEnumTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Oid:           OidColumn,
		Enumtypid:     EnumtypidColumn,
		Enumsortorder: EnumsortorderColumn,
		Enumlabel:     EnumlabelColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

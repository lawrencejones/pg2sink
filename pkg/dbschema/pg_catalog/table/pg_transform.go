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

var PgTransform = newPgTransformTable()

type PgTransformTable struct {
	postgres.Table

	//Columns
	Oid        postgres.ColumnString
	Trftype    postgres.ColumnString
	Trflang    postgres.ColumnString
	Trffromsql postgres.ColumnString
	Trftosql   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgTransformTable with assigned alias
func (a *PgTransformTable) AS(alias string) *PgTransformTable {
	aliasTable := newPgTransformTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgTransformTable() *PgTransformTable {
	var (
		OidColumn        = postgres.StringColumn("oid")
		TrftypeColumn    = postgres.StringColumn("trftype")
		TrflangColumn    = postgres.StringColumn("trflang")
		TrffromsqlColumn = postgres.StringColumn("trffromsql")
		TrftosqlColumn   = postgres.StringColumn("trftosql")
	)

	return &PgTransformTable{
		Table: postgres.NewTable("pg_catalog", "pg_transform", OidColumn, TrftypeColumn, TrflangColumn, TrffromsqlColumn, TrftosqlColumn),

		//Columns
		Oid:        OidColumn,
		Trftype:    TrftypeColumn,
		Trflang:    TrflangColumn,
		Trffromsql: TrffromsqlColumn,
		Trftosql:   TrftosqlColumn,

		AllColumns:     postgres.ColumnList{OidColumn, TrftypeColumn, TrflangColumn, TrffromsqlColumn, TrftosqlColumn},
		MutableColumns: postgres.ColumnList{OidColumn, TrftypeColumn, TrflangColumn, TrffromsqlColumn, TrftosqlColumn},
	}
}

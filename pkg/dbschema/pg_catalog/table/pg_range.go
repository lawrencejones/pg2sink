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

var PgRange = newPgRangeTable()

type PgRangeTable struct {
	postgres.Table

	//Columns
	Rngtypid     postgres.ColumnString
	Rngsubtype   postgres.ColumnString
	Rngcollation postgres.ColumnString
	Rngsubopc    postgres.ColumnString
	Rngcanonical postgres.ColumnString
	Rngsubdiff   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgRangeTable with assigned alias
func (a *PgRangeTable) AS(alias string) *PgRangeTable {
	aliasTable := newPgRangeTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgRangeTable() *PgRangeTable {
	var (
		RngtypidColumn     = postgres.StringColumn("rngtypid")
		RngsubtypeColumn   = postgres.StringColumn("rngsubtype")
		RngcollationColumn = postgres.StringColumn("rngcollation")
		RngsubopcColumn    = postgres.StringColumn("rngsubopc")
		RngcanonicalColumn = postgres.StringColumn("rngcanonical")
		RngsubdiffColumn   = postgres.StringColumn("rngsubdiff")
	)

	return &PgRangeTable{
		Table: postgres.NewTable("pg_catalog", "pg_range", RngtypidColumn, RngsubtypeColumn, RngcollationColumn, RngsubopcColumn, RngcanonicalColumn, RngsubdiffColumn),

		//Columns
		Rngtypid:     RngtypidColumn,
		Rngsubtype:   RngsubtypeColumn,
		Rngcollation: RngcollationColumn,
		Rngsubopc:    RngsubopcColumn,
		Rngcanonical: RngcanonicalColumn,
		Rngsubdiff:   RngsubdiffColumn,

		AllColumns:     postgres.ColumnList{RngtypidColumn, RngsubtypeColumn, RngcollationColumn, RngsubopcColumn, RngcanonicalColumn, RngsubdiffColumn},
		MutableColumns: postgres.ColumnList{RngtypidColumn, RngsubtypeColumn, RngcollationColumn, RngsubopcColumn, RngcanonicalColumn, RngsubdiffColumn},
	}
}

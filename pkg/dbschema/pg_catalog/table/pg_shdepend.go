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

var PgShdepend = newPgShdependTable()

type PgShdependTable struct {
	postgres.Table

	//Columns
	Dbid       postgres.ColumnString
	Classid    postgres.ColumnString
	Objid      postgres.ColumnString
	Objsubid   postgres.ColumnInteger
	Refclassid postgres.ColumnString
	Refobjid   postgres.ColumnString
	Deptype    postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgShdependTable with assigned alias
func (a *PgShdependTable) AS(alias string) *PgShdependTable {
	aliasTable := newPgShdependTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgShdependTable() *PgShdependTable {
	var (
		DbidColumn       = postgres.StringColumn("dbid")
		ClassidColumn    = postgres.StringColumn("classid")
		ObjidColumn      = postgres.StringColumn("objid")
		ObjsubidColumn   = postgres.IntegerColumn("objsubid")
		RefclassidColumn = postgres.StringColumn("refclassid")
		RefobjidColumn   = postgres.StringColumn("refobjid")
		DeptypeColumn    = postgres.StringColumn("deptype")
	)

	return &PgShdependTable{
		Table: postgres.NewTable("pg_catalog", "pg_shdepend", DbidColumn, ClassidColumn, ObjidColumn, ObjsubidColumn, RefclassidColumn, RefobjidColumn, DeptypeColumn),

		//Columns
		Dbid:       DbidColumn,
		Classid:    ClassidColumn,
		Objid:      ObjidColumn,
		Objsubid:   ObjsubidColumn,
		Refclassid: RefclassidColumn,
		Refobjid:   RefobjidColumn,
		Deptype:    DeptypeColumn,

		AllColumns:     postgres.ColumnList{DbidColumn, ClassidColumn, ObjidColumn, ObjsubidColumn, RefclassidColumn, RefobjidColumn, DeptypeColumn},
		MutableColumns: postgres.ColumnList{DbidColumn, ClassidColumn, ObjidColumn, ObjsubidColumn, RefclassidColumn, RefobjidColumn, DeptypeColumn},
	}
}

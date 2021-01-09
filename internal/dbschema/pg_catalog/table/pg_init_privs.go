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

var PgInitPrivs = newPgInitPrivsTable()

type PgInitPrivsTable struct {
	postgres.Table

	//Columns
	Objoid    postgres.ColumnString
	Classoid  postgres.ColumnString
	Objsubid  postgres.ColumnInteger
	Privtype  postgres.ColumnString
	Initprivs postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgInitPrivsTable with assigned alias
func (a *PgInitPrivsTable) AS(alias string) *PgInitPrivsTable {
	aliasTable := newPgInitPrivsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgInitPrivsTable() *PgInitPrivsTable {
	var (
		ObjoidColumn    = postgres.StringColumn("objoid")
		ClassoidColumn  = postgres.StringColumn("classoid")
		ObjsubidColumn  = postgres.IntegerColumn("objsubid")
		PrivtypeColumn  = postgres.StringColumn("privtype")
		InitprivsColumn = postgres.StringColumn("initprivs")
	)

	return &PgInitPrivsTable{
		Table: postgres.NewTable("pg_catalog", "pg_init_privs", ObjoidColumn, ClassoidColumn, ObjsubidColumn, PrivtypeColumn, InitprivsColumn),

		//Columns
		Objoid:    ObjoidColumn,
		Classoid:  ClassoidColumn,
		Objsubid:  ObjsubidColumn,
		Privtype:  PrivtypeColumn,
		Initprivs: InitprivsColumn,

		AllColumns:     postgres.ColumnList{ObjoidColumn, ClassoidColumn, ObjsubidColumn, PrivtypeColumn, InitprivsColumn},
		MutableColumns: postgres.ColumnList{ObjoidColumn, ClassoidColumn, ObjsubidColumn, PrivtypeColumn, InitprivsColumn},
	}
}
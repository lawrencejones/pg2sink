//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/postgres"
)

var PgStatUserFunctions = newPgStatUserFunctionsTable()

type PgStatUserFunctionsTable struct {
	postgres.Table

	//Columns
	Funcid     postgres.ColumnString
	Schemaname postgres.ColumnString
	Funcname   postgres.ColumnString
	Calls      postgres.ColumnInteger
	TotalTime  postgres.ColumnFloat
	SelfTime   postgres.ColumnFloat

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatUserFunctionsTable with assigned alias
func (a *PgStatUserFunctionsTable) AS(alias string) *PgStatUserFunctionsTable {
	aliasTable := newPgStatUserFunctionsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatUserFunctionsTable() *PgStatUserFunctionsTable {
	var (
		FuncidColumn     = postgres.StringColumn("funcid")
		SchemanameColumn = postgres.StringColumn("schemaname")
		FuncnameColumn   = postgres.StringColumn("funcname")
		CallsColumn      = postgres.IntegerColumn("calls")
		TotalTimeColumn  = postgres.FloatColumn("total_time")
		SelfTimeColumn   = postgres.FloatColumn("self_time")
	)

	return &PgStatUserFunctionsTable{
		Table: postgres.NewTable("pg_catalog", "pg_stat_user_functions", FuncidColumn, SchemanameColumn, FuncnameColumn, CallsColumn, TotalTimeColumn, SelfTimeColumn),

		//Columns
		Funcid:     FuncidColumn,
		Schemaname: SchemanameColumn,
		Funcname:   FuncnameColumn,
		Calls:      CallsColumn,
		TotalTime:  TotalTimeColumn,
		SelfTime:   SelfTimeColumn,

		AllColumns:     postgres.ColumnList{FuncidColumn, SchemanameColumn, FuncnameColumn, CallsColumn, TotalTimeColumn, SelfTimeColumn},
		MutableColumns: postgres.ColumnList{FuncidColumn, SchemanameColumn, FuncnameColumn, CallsColumn, TotalTimeColumn, SelfTimeColumn},
	}
}

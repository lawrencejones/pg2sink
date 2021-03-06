//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package view

import (
	"github.com/go-jet/jet/v2/postgres"
)

var PgUserMappings = newPgUserMappingsTable()

type pgUserMappingsTable struct {
	postgres.Table

	//Columns
	Umid      postgres.ColumnString
	Srvid     postgres.ColumnString
	Srvname   postgres.ColumnString
	Umuser    postgres.ColumnString
	Usename   postgres.ColumnString
	Umoptions postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgUserMappingsTable struct {
	pgUserMappingsTable

	EXCLUDED pgUserMappingsTable
}

// AS creates new PgUserMappingsTable with assigned alias
func (a *PgUserMappingsTable) AS(alias string) *PgUserMappingsTable {
	aliasTable := newPgUserMappingsTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgUserMappingsTable() *PgUserMappingsTable {
	return &PgUserMappingsTable{
		pgUserMappingsTable: newPgUserMappingsTableImpl("pg_catalog", "pg_user_mappings"),
		EXCLUDED:            newPgUserMappingsTableImpl("", "excluded"),
	}
}

func newPgUserMappingsTableImpl(schemaName, tableName string) pgUserMappingsTable {
	var (
		UmidColumn      = postgres.StringColumn("umid")
		SrvidColumn     = postgres.StringColumn("srvid")
		SrvnameColumn   = postgres.StringColumn("srvname")
		UmuserColumn    = postgres.StringColumn("umuser")
		UsenameColumn   = postgres.StringColumn("usename")
		UmoptionsColumn = postgres.StringColumn("umoptions")
		allColumns      = postgres.ColumnList{UmidColumn, SrvidColumn, SrvnameColumn, UmuserColumn, UsenameColumn, UmoptionsColumn}
		mutableColumns  = postgres.ColumnList{UmidColumn, SrvidColumn, SrvnameColumn, UmuserColumn, UsenameColumn, UmoptionsColumn}
	)

	return pgUserMappingsTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Umid:      UmidColumn,
		Srvid:     SrvidColumn,
		Srvname:   SrvnameColumn,
		Umuser:    UmuserColumn,
		Usename:   UsenameColumn,
		Umoptions: UmoptionsColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

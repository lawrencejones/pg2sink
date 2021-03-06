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

var PgDbRoleSetting = newPgDbRoleSettingTable()

type pgDbRoleSettingTable struct {
	postgres.Table

	//Columns
	Setdatabase postgres.ColumnString
	Setrole     postgres.ColumnString
	Setconfig   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgDbRoleSettingTable struct {
	pgDbRoleSettingTable

	EXCLUDED pgDbRoleSettingTable
}

// AS creates new PgDbRoleSettingTable with assigned alias
func (a *PgDbRoleSettingTable) AS(alias string) *PgDbRoleSettingTable {
	aliasTable := newPgDbRoleSettingTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgDbRoleSettingTable() *PgDbRoleSettingTable {
	return &PgDbRoleSettingTable{
		pgDbRoleSettingTable: newPgDbRoleSettingTableImpl("pg_catalog", "pg_db_role_setting"),
		EXCLUDED:             newPgDbRoleSettingTableImpl("", "excluded"),
	}
}

func newPgDbRoleSettingTableImpl(schemaName, tableName string) pgDbRoleSettingTable {
	var (
		SetdatabaseColumn = postgres.StringColumn("setdatabase")
		SetroleColumn     = postgres.StringColumn("setrole")
		SetconfigColumn   = postgres.StringColumn("setconfig")
		allColumns        = postgres.ColumnList{SetdatabaseColumn, SetroleColumn, SetconfigColumn}
		mutableColumns    = postgres.ColumnList{SetdatabaseColumn, SetroleColumn, SetconfigColumn}
	)

	return pgDbRoleSettingTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Setdatabase: SetdatabaseColumn,
		Setrole:     SetroleColumn,
		Setconfig:   SetconfigColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

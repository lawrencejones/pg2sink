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

var PgRoles = newPgRolesTable()

type pgRolesTable struct {
	postgres.Table

	//Columns
	Rolname        postgres.ColumnString
	Rolsuper       postgres.ColumnBool
	Rolinherit     postgres.ColumnBool
	Rolcreaterole  postgres.ColumnBool
	Rolcreatedb    postgres.ColumnBool
	Rolcanlogin    postgres.ColumnBool
	Rolreplication postgres.ColumnBool
	Rolconnlimit   postgres.ColumnInteger
	Rolpassword    postgres.ColumnString
	Rolvaliduntil  postgres.ColumnTimestampz
	Rolbypassrls   postgres.ColumnBool
	Rolconfig      postgres.ColumnString
	Oid            postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PgRolesTable struct {
	pgRolesTable

	EXCLUDED pgRolesTable
}

// AS creates new PgRolesTable with assigned alias
func (a *PgRolesTable) AS(alias string) *PgRolesTable {
	aliasTable := newPgRolesTable()
	aliasTable.Table.AS(alias)
	return aliasTable
}

func newPgRolesTable() *PgRolesTable {
	return &PgRolesTable{
		pgRolesTable: newPgRolesTableImpl("pg_catalog", "pg_roles"),
		EXCLUDED:     newPgRolesTableImpl("", "excluded"),
	}
}

func newPgRolesTableImpl(schemaName, tableName string) pgRolesTable {
	var (
		RolnameColumn        = postgres.StringColumn("rolname")
		RolsuperColumn       = postgres.BoolColumn("rolsuper")
		RolinheritColumn     = postgres.BoolColumn("rolinherit")
		RolcreateroleColumn  = postgres.BoolColumn("rolcreaterole")
		RolcreatedbColumn    = postgres.BoolColumn("rolcreatedb")
		RolcanloginColumn    = postgres.BoolColumn("rolcanlogin")
		RolreplicationColumn = postgres.BoolColumn("rolreplication")
		RolconnlimitColumn   = postgres.IntegerColumn("rolconnlimit")
		RolpasswordColumn    = postgres.StringColumn("rolpassword")
		RolvaliduntilColumn  = postgres.TimestampzColumn("rolvaliduntil")
		RolbypassrlsColumn   = postgres.BoolColumn("rolbypassrls")
		RolconfigColumn      = postgres.StringColumn("rolconfig")
		OidColumn            = postgres.StringColumn("oid")
		allColumns           = postgres.ColumnList{RolnameColumn, RolsuperColumn, RolinheritColumn, RolcreateroleColumn, RolcreatedbColumn, RolcanloginColumn, RolreplicationColumn, RolconnlimitColumn, RolpasswordColumn, RolvaliduntilColumn, RolbypassrlsColumn, RolconfigColumn, OidColumn}
		mutableColumns       = postgres.ColumnList{RolnameColumn, RolsuperColumn, RolinheritColumn, RolcreateroleColumn, RolcreatedbColumn, RolcanloginColumn, RolreplicationColumn, RolconnlimitColumn, RolpasswordColumn, RolvaliduntilColumn, RolbypassrlsColumn, RolconfigColumn, OidColumn}
	)

	return pgRolesTable{
		Table: postgres.NewTable(schemaName, tableName, allColumns...),

		//Columns
		Rolname:        RolnameColumn,
		Rolsuper:       RolsuperColumn,
		Rolinherit:     RolinheritColumn,
		Rolcreaterole:  RolcreateroleColumn,
		Rolcreatedb:    RolcreatedbColumn,
		Rolcanlogin:    RolcanloginColumn,
		Rolreplication: RolreplicationColumn,
		Rolconnlimit:   RolconnlimitColumn,
		Rolpassword:    RolpasswordColumn,
		Rolvaliduntil:  RolvaliduntilColumn,
		Rolbypassrls:   RolbypassrlsColumn,
		Rolconfig:      RolconfigColumn,
		Oid:            OidColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

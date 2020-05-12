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

var RoutinePrivileges = newRoutinePrivilegesTable()

type RoutinePrivilegesTable struct {
	postgres.Table

	//Columns
	Grantor         postgres.ColumnString
	Grantee         postgres.ColumnString
	SpecificCatalog postgres.ColumnString
	SpecificSchema  postgres.ColumnString
	SpecificName    postgres.ColumnString
	RoutineCatalog  postgres.ColumnString
	RoutineSchema   postgres.ColumnString
	RoutineName     postgres.ColumnString
	PrivilegeType   postgres.ColumnString
	IsGrantable     postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new RoutinePrivilegesTable with assigned alias
func (a *RoutinePrivilegesTable) AS(alias string) *RoutinePrivilegesTable {
	aliasTable := newRoutinePrivilegesTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newRoutinePrivilegesTable() *RoutinePrivilegesTable {
	var (
		GrantorColumn         = postgres.StringColumn("grantor")
		GranteeColumn         = postgres.StringColumn("grantee")
		SpecificCatalogColumn = postgres.StringColumn("specific_catalog")
		SpecificSchemaColumn  = postgres.StringColumn("specific_schema")
		SpecificNameColumn    = postgres.StringColumn("specific_name")
		RoutineCatalogColumn  = postgres.StringColumn("routine_catalog")
		RoutineSchemaColumn   = postgres.StringColumn("routine_schema")
		RoutineNameColumn     = postgres.StringColumn("routine_name")
		PrivilegeTypeColumn   = postgres.StringColumn("privilege_type")
		IsGrantableColumn     = postgres.StringColumn("is_grantable")
	)

	return &RoutinePrivilegesTable{
		Table: postgres.NewTable("information_schema", "routine_privileges", GrantorColumn, GranteeColumn, SpecificCatalogColumn, SpecificSchemaColumn, SpecificNameColumn, RoutineCatalogColumn, RoutineSchemaColumn, RoutineNameColumn, PrivilegeTypeColumn, IsGrantableColumn),

		//Columns
		Grantor:         GrantorColumn,
		Grantee:         GranteeColumn,
		SpecificCatalog: SpecificCatalogColumn,
		SpecificSchema:  SpecificSchemaColumn,
		SpecificName:    SpecificNameColumn,
		RoutineCatalog:  RoutineCatalogColumn,
		RoutineSchema:   RoutineSchemaColumn,
		RoutineName:     RoutineNameColumn,
		PrivilegeType:   PrivilegeTypeColumn,
		IsGrantable:     IsGrantableColumn,

		AllColumns:     postgres.ColumnList{GrantorColumn, GranteeColumn, SpecificCatalogColumn, SpecificSchemaColumn, SpecificNameColumn, RoutineCatalogColumn, RoutineSchemaColumn, RoutineNameColumn, PrivilegeTypeColumn, IsGrantableColumn},
		MutableColumns: postgres.ColumnList{GrantorColumn, GranteeColumn, SpecificCatalogColumn, SpecificSchemaColumn, SpecificNameColumn, RoutineCatalogColumn, RoutineSchemaColumn, RoutineNameColumn, PrivilegeTypeColumn, IsGrantableColumn},
	}
}

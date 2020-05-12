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

var RoleColumnGrants = newRoleColumnGrantsTable()

type RoleColumnGrantsTable struct {
	postgres.Table

	//Columns
	Grantor       postgres.ColumnString
	Grantee       postgres.ColumnString
	TableCatalog  postgres.ColumnString
	TableSchema   postgres.ColumnString
	TableName     postgres.ColumnString
	ColumnName    postgres.ColumnString
	PrivilegeType postgres.ColumnString
	IsGrantable   postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new RoleColumnGrantsTable with assigned alias
func (a *RoleColumnGrantsTable) AS(alias string) *RoleColumnGrantsTable {
	aliasTable := newRoleColumnGrantsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newRoleColumnGrantsTable() *RoleColumnGrantsTable {
	var (
		GrantorColumn       = postgres.StringColumn("grantor")
		GranteeColumn       = postgres.StringColumn("grantee")
		TableCatalogColumn  = postgres.StringColumn("table_catalog")
		TableSchemaColumn   = postgres.StringColumn("table_schema")
		TableNameColumn     = postgres.StringColumn("table_name")
		ColumnNameColumn    = postgres.StringColumn("column_name")
		PrivilegeTypeColumn = postgres.StringColumn("privilege_type")
		IsGrantableColumn   = postgres.StringColumn("is_grantable")
	)

	return &RoleColumnGrantsTable{
		Table: postgres.NewTable("information_schema", "role_column_grants", GrantorColumn, GranteeColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, PrivilegeTypeColumn, IsGrantableColumn),

		//Columns
		Grantor:       GrantorColumn,
		Grantee:       GranteeColumn,
		TableCatalog:  TableCatalogColumn,
		TableSchema:   TableSchemaColumn,
		TableName:     TableNameColumn,
		ColumnName:    ColumnNameColumn,
		PrivilegeType: PrivilegeTypeColumn,
		IsGrantable:   IsGrantableColumn,

		AllColumns:     postgres.ColumnList{GrantorColumn, GranteeColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, PrivilegeTypeColumn, IsGrantableColumn},
		MutableColumns: postgres.ColumnList{GrantorColumn, GranteeColumn, TableCatalogColumn, TableSchemaColumn, TableNameColumn, ColumnNameColumn, PrivilegeTypeColumn, IsGrantableColumn},
	}
}

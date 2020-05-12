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

var AdministrableRoleAuthorizations = newAdministrableRoleAuthorizationsTable()

type AdministrableRoleAuthorizationsTable struct {
	postgres.Table

	//Columns
	Grantee     postgres.ColumnString
	RoleName    postgres.ColumnString
	IsGrantable postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new AdministrableRoleAuthorizationsTable with assigned alias
func (a *AdministrableRoleAuthorizationsTable) AS(alias string) *AdministrableRoleAuthorizationsTable {
	aliasTable := newAdministrableRoleAuthorizationsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newAdministrableRoleAuthorizationsTable() *AdministrableRoleAuthorizationsTable {
	var (
		GranteeColumn     = postgres.StringColumn("grantee")
		RoleNameColumn    = postgres.StringColumn("role_name")
		IsGrantableColumn = postgres.StringColumn("is_grantable")
	)

	return &AdministrableRoleAuthorizationsTable{
		Table: postgres.NewTable("information_schema", "administrable_role_authorizations", GranteeColumn, RoleNameColumn, IsGrantableColumn),

		//Columns
		Grantee:     GranteeColumn,
		RoleName:    RoleNameColumn,
		IsGrantable: IsGrantableColumn,

		AllColumns:     postgres.ColumnList{GranteeColumn, RoleNameColumn, IsGrantableColumn},
		MutableColumns: postgres.ColumnList{GranteeColumn, RoleNameColumn, IsGrantableColumn},
	}
}

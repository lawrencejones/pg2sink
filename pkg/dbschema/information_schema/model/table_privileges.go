//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 09:15:06 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type TablePrivileges struct {
	Grantor       *string
	Grantee       *string
	TableCatalog  *string
	TableSchema   *string
	TableName     *string
	PrivilegeType *string
	IsGrantable   *string
	WithHierarchy *string
}

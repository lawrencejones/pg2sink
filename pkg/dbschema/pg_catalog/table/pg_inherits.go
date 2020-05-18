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

var PgInherits = newPgInheritsTable()

type PgInheritsTable struct {
	postgres.Table

	//Columns
	Inhrelid  postgres.ColumnString
	Inhparent postgres.ColumnString
	Inhseqno  postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgInheritsTable with assigned alias
func (a *PgInheritsTable) AS(alias string) *PgInheritsTable {
	aliasTable := newPgInheritsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgInheritsTable() *PgInheritsTable {
	var (
		InhrelidColumn  = postgres.StringColumn("inhrelid")
		InhparentColumn = postgres.StringColumn("inhparent")
		InhseqnoColumn  = postgres.IntegerColumn("inhseqno")
	)

	return &PgInheritsTable{
		Table: postgres.NewTable("pg_catalog", "pg_inherits", InhrelidColumn, InhparentColumn, InhseqnoColumn),

		//Columns
		Inhrelid:  InhrelidColumn,
		Inhparent: InhparentColumn,
		Inhseqno:  InhseqnoColumn,

		AllColumns:     postgres.ColumnList{InhrelidColumn, InhparentColumn, InhseqnoColumn},
		MutableColumns: postgres.ColumnList{InhrelidColumn, InhparentColumn, InhseqnoColumn},
	}
}

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

var PgPltemplate = newPgPltemplateTable()

type PgPltemplateTable struct {
	postgres.Table

	//Columns
	Tmplname      postgres.ColumnString
	Tmpltrusted   postgres.ColumnBool
	Tmpldbacreate postgres.ColumnBool
	Tmplhandler   postgres.ColumnString
	Tmplinline    postgres.ColumnString
	Tmplvalidator postgres.ColumnString
	Tmpllibrary   postgres.ColumnString
	Tmplacl       postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgPltemplateTable with assigned alias
func (a *PgPltemplateTable) AS(alias string) *PgPltemplateTable {
	aliasTable := newPgPltemplateTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgPltemplateTable() *PgPltemplateTable {
	var (
		TmplnameColumn      = postgres.StringColumn("tmplname")
		TmpltrustedColumn   = postgres.BoolColumn("tmpltrusted")
		TmpldbacreateColumn = postgres.BoolColumn("tmpldbacreate")
		TmplhandlerColumn   = postgres.StringColumn("tmplhandler")
		TmplinlineColumn    = postgres.StringColumn("tmplinline")
		TmplvalidatorColumn = postgres.StringColumn("tmplvalidator")
		TmpllibraryColumn   = postgres.StringColumn("tmpllibrary")
		TmplaclColumn       = postgres.StringColumn("tmplacl")
	)

	return &PgPltemplateTable{
		Table: postgres.NewTable("pg_catalog", "pg_pltemplate", TmplnameColumn, TmpltrustedColumn, TmpldbacreateColumn, TmplhandlerColumn, TmplinlineColumn, TmplvalidatorColumn, TmpllibraryColumn, TmplaclColumn),

		//Columns
		Tmplname:      TmplnameColumn,
		Tmpltrusted:   TmpltrustedColumn,
		Tmpldbacreate: TmpldbacreateColumn,
		Tmplhandler:   TmplhandlerColumn,
		Tmplinline:    TmplinlineColumn,
		Tmplvalidator: TmplvalidatorColumn,
		Tmpllibrary:   TmpllibraryColumn,
		Tmplacl:       TmplaclColumn,

		AllColumns:     postgres.ColumnList{TmplnameColumn, TmpltrustedColumn, TmpldbacreateColumn, TmplhandlerColumn, TmplinlineColumn, TmplvalidatorColumn, TmpllibraryColumn, TmplaclColumn},
		MutableColumns: postgres.ColumnList{TmplnameColumn, TmpltrustedColumn, TmpldbacreateColumn, TmplhandlerColumn, TmplinlineColumn, TmplvalidatorColumn, TmpllibraryColumn, TmplaclColumn},
	}
}

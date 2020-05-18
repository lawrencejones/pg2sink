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

var PgStatioUserSequences = newPgStatioUserSequencesTable()

type PgStatioUserSequencesTable struct {
	postgres.Table

	//Columns
	Relid      postgres.ColumnString
	Schemaname postgres.ColumnString
	Relname    postgres.ColumnString
	BlksRead   postgres.ColumnInteger
	BlksHit    postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatioUserSequencesTable with assigned alias
func (a *PgStatioUserSequencesTable) AS(alias string) *PgStatioUserSequencesTable {
	aliasTable := newPgStatioUserSequencesTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatioUserSequencesTable() *PgStatioUserSequencesTable {
	var (
		RelidColumn      = postgres.StringColumn("relid")
		SchemanameColumn = postgres.StringColumn("schemaname")
		RelnameColumn    = postgres.StringColumn("relname")
		BlksReadColumn   = postgres.IntegerColumn("blks_read")
		BlksHitColumn    = postgres.IntegerColumn("blks_hit")
	)

	return &PgStatioUserSequencesTable{
		Table: postgres.NewTable("pg_catalog", "pg_statio_user_sequences", RelidColumn, SchemanameColumn, RelnameColumn, BlksReadColumn, BlksHitColumn),

		//Columns
		Relid:      RelidColumn,
		Schemaname: SchemanameColumn,
		Relname:    RelnameColumn,
		BlksRead:   BlksReadColumn,
		BlksHit:    BlksHitColumn,

		AllColumns:     postgres.ColumnList{RelidColumn, SchemanameColumn, RelnameColumn, BlksReadColumn, BlksHitColumn},
		MutableColumns: postgres.ColumnList{RelidColumn, SchemanameColumn, RelnameColumn, BlksReadColumn, BlksHitColumn},
	}
}

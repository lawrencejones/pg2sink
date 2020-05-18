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

var PgStatDatabaseConflicts = newPgStatDatabaseConflictsTable()

type PgStatDatabaseConflictsTable struct {
	postgres.Table

	//Columns
	Datid           postgres.ColumnString
	Datname         postgres.ColumnString
	ConflTablespace postgres.ColumnInteger
	ConflLock       postgres.ColumnInteger
	ConflSnapshot   postgres.ColumnInteger
	ConflBufferpin  postgres.ColumnInteger
	ConflDeadlock   postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatDatabaseConflictsTable with assigned alias
func (a *PgStatDatabaseConflictsTable) AS(alias string) *PgStatDatabaseConflictsTable {
	aliasTable := newPgStatDatabaseConflictsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatDatabaseConflictsTable() *PgStatDatabaseConflictsTable {
	var (
		DatidColumn           = postgres.StringColumn("datid")
		DatnameColumn         = postgres.StringColumn("datname")
		ConflTablespaceColumn = postgres.IntegerColumn("confl_tablespace")
		ConflLockColumn       = postgres.IntegerColumn("confl_lock")
		ConflSnapshotColumn   = postgres.IntegerColumn("confl_snapshot")
		ConflBufferpinColumn  = postgres.IntegerColumn("confl_bufferpin")
		ConflDeadlockColumn   = postgres.IntegerColumn("confl_deadlock")
	)

	return &PgStatDatabaseConflictsTable{
		Table: postgres.NewTable("pg_catalog", "pg_stat_database_conflicts", DatidColumn, DatnameColumn, ConflTablespaceColumn, ConflLockColumn, ConflSnapshotColumn, ConflBufferpinColumn, ConflDeadlockColumn),

		//Columns
		Datid:           DatidColumn,
		Datname:         DatnameColumn,
		ConflTablespace: ConflTablespaceColumn,
		ConflLock:       ConflLockColumn,
		ConflSnapshot:   ConflSnapshotColumn,
		ConflBufferpin:  ConflBufferpinColumn,
		ConflDeadlock:   ConflDeadlockColumn,

		AllColumns:     postgres.ColumnList{DatidColumn, DatnameColumn, ConflTablespaceColumn, ConflLockColumn, ConflSnapshotColumn, ConflBufferpinColumn, ConflDeadlockColumn},
		MutableColumns: postgres.ColumnList{DatidColumn, DatnameColumn, ConflTablespaceColumn, ConflLockColumn, ConflSnapshotColumn, ConflBufferpinColumn, ConflDeadlockColumn},
	}
}

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

var PgStatProgressVacuum = newPgStatProgressVacuumTable()

type PgStatProgressVacuumTable struct {
	postgres.Table

	//Columns
	Pid              postgres.ColumnInteger
	Datid            postgres.ColumnString
	Datname          postgres.ColumnString
	Relid            postgres.ColumnString
	Phase            postgres.ColumnString
	HeapBlksTotal    postgres.ColumnInteger
	HeapBlksScanned  postgres.ColumnInteger
	HeapBlksVacuumed postgres.ColumnInteger
	IndexVacuumCount postgres.ColumnInteger
	MaxDeadTuples    postgres.ColumnInteger
	NumDeadTuples    postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatProgressVacuumTable with assigned alias
func (a *PgStatProgressVacuumTable) AS(alias string) *PgStatProgressVacuumTable {
	aliasTable := newPgStatProgressVacuumTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatProgressVacuumTable() *PgStatProgressVacuumTable {
	var (
		PidColumn              = postgres.IntegerColumn("pid")
		DatidColumn            = postgres.StringColumn("datid")
		DatnameColumn          = postgres.StringColumn("datname")
		RelidColumn            = postgres.StringColumn("relid")
		PhaseColumn            = postgres.StringColumn("phase")
		HeapBlksTotalColumn    = postgres.IntegerColumn("heap_blks_total")
		HeapBlksScannedColumn  = postgres.IntegerColumn("heap_blks_scanned")
		HeapBlksVacuumedColumn = postgres.IntegerColumn("heap_blks_vacuumed")
		IndexVacuumCountColumn = postgres.IntegerColumn("index_vacuum_count")
		MaxDeadTuplesColumn    = postgres.IntegerColumn("max_dead_tuples")
		NumDeadTuplesColumn    = postgres.IntegerColumn("num_dead_tuples")
	)

	return &PgStatProgressVacuumTable{
		Table: postgres.NewTable("pg_catalog", "pg_stat_progress_vacuum", PidColumn, DatidColumn, DatnameColumn, RelidColumn, PhaseColumn, HeapBlksTotalColumn, HeapBlksScannedColumn, HeapBlksVacuumedColumn, IndexVacuumCountColumn, MaxDeadTuplesColumn, NumDeadTuplesColumn),

		//Columns
		Pid:              PidColumn,
		Datid:            DatidColumn,
		Datname:          DatnameColumn,
		Relid:            RelidColumn,
		Phase:            PhaseColumn,
		HeapBlksTotal:    HeapBlksTotalColumn,
		HeapBlksScanned:  HeapBlksScannedColumn,
		HeapBlksVacuumed: HeapBlksVacuumedColumn,
		IndexVacuumCount: IndexVacuumCountColumn,
		MaxDeadTuples:    MaxDeadTuplesColumn,
		NumDeadTuples:    NumDeadTuplesColumn,

		AllColumns:     postgres.ColumnList{PidColumn, DatidColumn, DatnameColumn, RelidColumn, PhaseColumn, HeapBlksTotalColumn, HeapBlksScannedColumn, HeapBlksVacuumedColumn, IndexVacuumCountColumn, MaxDeadTuplesColumn, NumDeadTuplesColumn},
		MutableColumns: postgres.ColumnList{PidColumn, DatidColumn, DatnameColumn, RelidColumn, PhaseColumn, HeapBlksTotalColumn, HeapBlksScannedColumn, HeapBlksVacuumedColumn, IndexVacuumCountColumn, MaxDeadTuplesColumn, NumDeadTuplesColumn},
	}
}
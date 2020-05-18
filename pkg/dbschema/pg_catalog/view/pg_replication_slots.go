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

var PgReplicationSlots = newPgReplicationSlotsTable()

type PgReplicationSlotsTable struct {
	postgres.Table

	//Columns
	SlotName          postgres.ColumnString
	Plugin            postgres.ColumnString
	SlotType          postgres.ColumnString
	Datoid            postgres.ColumnString
	Database          postgres.ColumnString
	Temporary         postgres.ColumnBool
	Active            postgres.ColumnBool
	ActivePid         postgres.ColumnInteger
	Xmin              postgres.ColumnString
	CatalogXmin       postgres.ColumnString
	RestartLsn        postgres.ColumnString
	ConfirmedFlushLsn postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgReplicationSlotsTable with assigned alias
func (a *PgReplicationSlotsTable) AS(alias string) *PgReplicationSlotsTable {
	aliasTable := newPgReplicationSlotsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgReplicationSlotsTable() *PgReplicationSlotsTable {
	var (
		SlotNameColumn          = postgres.StringColumn("slot_name")
		PluginColumn            = postgres.StringColumn("plugin")
		SlotTypeColumn          = postgres.StringColumn("slot_type")
		DatoidColumn            = postgres.StringColumn("datoid")
		DatabaseColumn          = postgres.StringColumn("database")
		TemporaryColumn         = postgres.BoolColumn("temporary")
		ActiveColumn            = postgres.BoolColumn("active")
		ActivePidColumn         = postgres.IntegerColumn("active_pid")
		XminColumn              = postgres.StringColumn("xmin")
		CatalogXminColumn       = postgres.StringColumn("catalog_xmin")
		RestartLsnColumn        = postgres.StringColumn("restart_lsn")
		ConfirmedFlushLsnColumn = postgres.StringColumn("confirmed_flush_lsn")
	)

	return &PgReplicationSlotsTable{
		Table: postgres.NewTable("pg_catalog", "pg_replication_slots", SlotNameColumn, PluginColumn, SlotTypeColumn, DatoidColumn, DatabaseColumn, TemporaryColumn, ActiveColumn, ActivePidColumn, XminColumn, CatalogXminColumn, RestartLsnColumn, ConfirmedFlushLsnColumn),

		//Columns
		SlotName:          SlotNameColumn,
		Plugin:            PluginColumn,
		SlotType:          SlotTypeColumn,
		Datoid:            DatoidColumn,
		Database:          DatabaseColumn,
		Temporary:         TemporaryColumn,
		Active:            ActiveColumn,
		ActivePid:         ActivePidColumn,
		Xmin:              XminColumn,
		CatalogXmin:       CatalogXminColumn,
		RestartLsn:        RestartLsnColumn,
		ConfirmedFlushLsn: ConfirmedFlushLsnColumn,

		AllColumns:     postgres.ColumnList{SlotNameColumn, PluginColumn, SlotTypeColumn, DatoidColumn, DatabaseColumn, TemporaryColumn, ActiveColumn, ActivePidColumn, XminColumn, CatalogXminColumn, RestartLsnColumn, ConfirmedFlushLsnColumn},
		MutableColumns: postgres.ColumnList{SlotNameColumn, PluginColumn, SlotTypeColumn, DatoidColumn, DatabaseColumn, TemporaryColumn, ActiveColumn, ActivePidColumn, XminColumn, CatalogXminColumn, RestartLsnColumn, ConfirmedFlushLsnColumn},
	}
}

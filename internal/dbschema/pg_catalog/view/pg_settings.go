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

var PgSettings = newPgSettingsTable()

type PgSettingsTable struct {
	postgres.Table

	//Columns
	Name           postgres.ColumnString
	Setting        postgres.ColumnString
	Unit           postgres.ColumnString
	Category       postgres.ColumnString
	ShortDesc      postgres.ColumnString
	ExtraDesc      postgres.ColumnString
	Context        postgres.ColumnString
	Vartype        postgres.ColumnString
	Source         postgres.ColumnString
	MinVal         postgres.ColumnString
	MaxVal         postgres.ColumnString
	Enumvals       postgres.ColumnString
	BootVal        postgres.ColumnString
	ResetVal       postgres.ColumnString
	Sourcefile     postgres.ColumnString
	Sourceline     postgres.ColumnInteger
	PendingRestart postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgSettingsTable with assigned alias
func (a *PgSettingsTable) AS(alias string) *PgSettingsTable {
	aliasTable := newPgSettingsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgSettingsTable() *PgSettingsTable {
	var (
		NameColumn           = postgres.StringColumn("name")
		SettingColumn        = postgres.StringColumn("setting")
		UnitColumn           = postgres.StringColumn("unit")
		CategoryColumn       = postgres.StringColumn("category")
		ShortDescColumn      = postgres.StringColumn("short_desc")
		ExtraDescColumn      = postgres.StringColumn("extra_desc")
		ContextColumn        = postgres.StringColumn("context")
		VartypeColumn        = postgres.StringColumn("vartype")
		SourceColumn         = postgres.StringColumn("source")
		MinValColumn         = postgres.StringColumn("min_val")
		MaxValColumn         = postgres.StringColumn("max_val")
		EnumvalsColumn       = postgres.StringColumn("enumvals")
		BootValColumn        = postgres.StringColumn("boot_val")
		ResetValColumn       = postgres.StringColumn("reset_val")
		SourcefileColumn     = postgres.StringColumn("sourcefile")
		SourcelineColumn     = postgres.IntegerColumn("sourceline")
		PendingRestartColumn = postgres.BoolColumn("pending_restart")
	)

	return &PgSettingsTable{
		Table: postgres.NewTable("pg_catalog", "pg_settings", NameColumn, SettingColumn, UnitColumn, CategoryColumn, ShortDescColumn, ExtraDescColumn, ContextColumn, VartypeColumn, SourceColumn, MinValColumn, MaxValColumn, EnumvalsColumn, BootValColumn, ResetValColumn, SourcefileColumn, SourcelineColumn, PendingRestartColumn),

		//Columns
		Name:           NameColumn,
		Setting:        SettingColumn,
		Unit:           UnitColumn,
		Category:       CategoryColumn,
		ShortDesc:      ShortDescColumn,
		ExtraDesc:      ExtraDescColumn,
		Context:        ContextColumn,
		Vartype:        VartypeColumn,
		Source:         SourceColumn,
		MinVal:         MinValColumn,
		MaxVal:         MaxValColumn,
		Enumvals:       EnumvalsColumn,
		BootVal:        BootValColumn,
		ResetVal:       ResetValColumn,
		Sourcefile:     SourcefileColumn,
		Sourceline:     SourcelineColumn,
		PendingRestart: PendingRestartColumn,

		AllColumns:     postgres.ColumnList{NameColumn, SettingColumn, UnitColumn, CategoryColumn, ShortDescColumn, ExtraDescColumn, ContextColumn, VartypeColumn, SourceColumn, MinValColumn, MaxValColumn, EnumvalsColumn, BootValColumn, ResetValColumn, SourcefileColumn, SourcelineColumn, PendingRestartColumn},
		MutableColumns: postgres.ColumnList{NameColumn, SettingColumn, UnitColumn, CategoryColumn, ShortDescColumn, ExtraDescColumn, ContextColumn, VartypeColumn, SourceColumn, MinValColumn, MaxValColumn, EnumvalsColumn, BootValColumn, ResetValColumn, SourcefileColumn, SourcelineColumn, PendingRestartColumn},
	}
}
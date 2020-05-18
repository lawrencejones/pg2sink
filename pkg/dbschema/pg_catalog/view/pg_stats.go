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

var PgStats = newPgStatsTable()

type PgStatsTable struct {
	postgres.Table

	//Columns
	Schemaname          postgres.ColumnString
	Tablename           postgres.ColumnString
	Attname             postgres.ColumnString
	Inherited           postgres.ColumnBool
	NullFrac            postgres.ColumnFloat
	AvgWidth            postgres.ColumnInteger
	NDistinct           postgres.ColumnFloat
	MostCommonVals      postgres.ColumnString
	MostCommonFreqs     postgres.ColumnString
	HistogramBounds     postgres.ColumnString
	Correlation         postgres.ColumnFloat
	MostCommonElems     postgres.ColumnString
	MostCommonElemFreqs postgres.ColumnString
	ElemCountHistogram  postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgStatsTable with assigned alias
func (a *PgStatsTable) AS(alias string) *PgStatsTable {
	aliasTable := newPgStatsTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgStatsTable() *PgStatsTable {
	var (
		SchemanameColumn          = postgres.StringColumn("schemaname")
		TablenameColumn           = postgres.StringColumn("tablename")
		AttnameColumn             = postgres.StringColumn("attname")
		InheritedColumn           = postgres.BoolColumn("inherited")
		NullFracColumn            = postgres.FloatColumn("null_frac")
		AvgWidthColumn            = postgres.IntegerColumn("avg_width")
		NDistinctColumn           = postgres.FloatColumn("n_distinct")
		MostCommonValsColumn      = postgres.StringColumn("most_common_vals")
		MostCommonFreqsColumn     = postgres.StringColumn("most_common_freqs")
		HistogramBoundsColumn     = postgres.StringColumn("histogram_bounds")
		CorrelationColumn         = postgres.FloatColumn("correlation")
		MostCommonElemsColumn     = postgres.StringColumn("most_common_elems")
		MostCommonElemFreqsColumn = postgres.StringColumn("most_common_elem_freqs")
		ElemCountHistogramColumn  = postgres.StringColumn("elem_count_histogram")
	)

	return &PgStatsTable{
		Table: postgres.NewTable("pg_catalog", "pg_stats", SchemanameColumn, TablenameColumn, AttnameColumn, InheritedColumn, NullFracColumn, AvgWidthColumn, NDistinctColumn, MostCommonValsColumn, MostCommonFreqsColumn, HistogramBoundsColumn, CorrelationColumn, MostCommonElemsColumn, MostCommonElemFreqsColumn, ElemCountHistogramColumn),

		//Columns
		Schemaname:          SchemanameColumn,
		Tablename:           TablenameColumn,
		Attname:             AttnameColumn,
		Inherited:           InheritedColumn,
		NullFrac:            NullFracColumn,
		AvgWidth:            AvgWidthColumn,
		NDistinct:           NDistinctColumn,
		MostCommonVals:      MostCommonValsColumn,
		MostCommonFreqs:     MostCommonFreqsColumn,
		HistogramBounds:     HistogramBoundsColumn,
		Correlation:         CorrelationColumn,
		MostCommonElems:     MostCommonElemsColumn,
		MostCommonElemFreqs: MostCommonElemFreqsColumn,
		ElemCountHistogram:  ElemCountHistogramColumn,

		AllColumns:     postgres.ColumnList{SchemanameColumn, TablenameColumn, AttnameColumn, InheritedColumn, NullFracColumn, AvgWidthColumn, NDistinctColumn, MostCommonValsColumn, MostCommonFreqsColumn, HistogramBoundsColumn, CorrelationColumn, MostCommonElemsColumn, MostCommonElemFreqsColumn, ElemCountHistogramColumn},
		MutableColumns: postgres.ColumnList{SchemanameColumn, TablenameColumn, AttnameColumn, InheritedColumn, NullFracColumn, AvgWidthColumn, NDistinctColumn, MostCommonValsColumn, MostCommonFreqsColumn, HistogramBoundsColumn, CorrelationColumn, MostCommonElemsColumn, MostCommonElemFreqsColumn, ElemCountHistogramColumn},
	}
}

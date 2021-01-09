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

var PgIndex = newPgIndexTable()

type PgIndexTable struct {
	postgres.Table

	//Columns
	Indexrelid     postgres.ColumnString
	Indrelid       postgres.ColumnString
	Indnatts       postgres.ColumnInteger
	Indnkeyatts    postgres.ColumnInteger
	Indisunique    postgres.ColumnBool
	Indisprimary   postgres.ColumnBool
	Indisexclusion postgres.ColumnBool
	Indimmediate   postgres.ColumnBool
	Indisclustered postgres.ColumnBool
	Indisvalid     postgres.ColumnBool
	Indcheckxmin   postgres.ColumnBool
	Indisready     postgres.ColumnBool
	Indislive      postgres.ColumnBool
	Indisreplident postgres.ColumnBool
	Indkey         postgres.ColumnString
	Indcollation   postgres.ColumnString
	Indclass       postgres.ColumnString
	Indoption      postgres.ColumnString
	Indexprs       postgres.ColumnString
	Indpred        postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

// creates new PgIndexTable with assigned alias
func (a *PgIndexTable) AS(alias string) *PgIndexTable {
	aliasTable := newPgIndexTable()

	aliasTable.Table.AS(alias)

	return aliasTable
}

func newPgIndexTable() *PgIndexTable {
	var (
		IndexrelidColumn     = postgres.StringColumn("indexrelid")
		IndrelidColumn       = postgres.StringColumn("indrelid")
		IndnattsColumn       = postgres.IntegerColumn("indnatts")
		IndnkeyattsColumn    = postgres.IntegerColumn("indnkeyatts")
		IndisuniqueColumn    = postgres.BoolColumn("indisunique")
		IndisprimaryColumn   = postgres.BoolColumn("indisprimary")
		IndisexclusionColumn = postgres.BoolColumn("indisexclusion")
		IndimmediateColumn   = postgres.BoolColumn("indimmediate")
		IndisclusteredColumn = postgres.BoolColumn("indisclustered")
		IndisvalidColumn     = postgres.BoolColumn("indisvalid")
		IndcheckxminColumn   = postgres.BoolColumn("indcheckxmin")
		IndisreadyColumn     = postgres.BoolColumn("indisready")
		IndisliveColumn      = postgres.BoolColumn("indislive")
		IndisreplidentColumn = postgres.BoolColumn("indisreplident")
		IndkeyColumn         = postgres.StringColumn("indkey")
		IndcollationColumn   = postgres.StringColumn("indcollation")
		IndclassColumn       = postgres.StringColumn("indclass")
		IndoptionColumn      = postgres.StringColumn("indoption")
		IndexprsColumn       = postgres.StringColumn("indexprs")
		IndpredColumn        = postgres.StringColumn("indpred")
	)

	return &PgIndexTable{
		Table: postgres.NewTable("pg_catalog", "pg_index", IndexrelidColumn, IndrelidColumn, IndnattsColumn, IndnkeyattsColumn, IndisuniqueColumn, IndisprimaryColumn, IndisexclusionColumn, IndimmediateColumn, IndisclusteredColumn, IndisvalidColumn, IndcheckxminColumn, IndisreadyColumn, IndisliveColumn, IndisreplidentColumn, IndkeyColumn, IndcollationColumn, IndclassColumn, IndoptionColumn, IndexprsColumn, IndpredColumn),

		//Columns
		Indexrelid:     IndexrelidColumn,
		Indrelid:       IndrelidColumn,
		Indnatts:       IndnattsColumn,
		Indnkeyatts:    IndnkeyattsColumn,
		Indisunique:    IndisuniqueColumn,
		Indisprimary:   IndisprimaryColumn,
		Indisexclusion: IndisexclusionColumn,
		Indimmediate:   IndimmediateColumn,
		Indisclustered: IndisclusteredColumn,
		Indisvalid:     IndisvalidColumn,
		Indcheckxmin:   IndcheckxminColumn,
		Indisready:     IndisreadyColumn,
		Indislive:      IndisliveColumn,
		Indisreplident: IndisreplidentColumn,
		Indkey:         IndkeyColumn,
		Indcollation:   IndcollationColumn,
		Indclass:       IndclassColumn,
		Indoption:      IndoptionColumn,
		Indexprs:       IndexprsColumn,
		Indpred:        IndpredColumn,

		AllColumns:     postgres.ColumnList{IndexrelidColumn, IndrelidColumn, IndnattsColumn, IndnkeyattsColumn, IndisuniqueColumn, IndisprimaryColumn, IndisexclusionColumn, IndimmediateColumn, IndisclusteredColumn, IndisvalidColumn, IndcheckxminColumn, IndisreadyColumn, IndisliveColumn, IndisreplidentColumn, IndkeyColumn, IndcollationColumn, IndclassColumn, IndoptionColumn, IndexprsColumn, IndpredColumn},
		MutableColumns: postgres.ColumnList{IndexrelidColumn, IndrelidColumn, IndnattsColumn, IndnkeyattsColumn, IndisuniqueColumn, IndisprimaryColumn, IndisexclusionColumn, IndimmediateColumn, IndisclusteredColumn, IndisvalidColumn, IndcheckxminColumn, IndisreadyColumn, IndisliveColumn, IndisreplidentColumn, IndkeyColumn, IndcollationColumn, IndclassColumn, IndoptionColumn, IndexprsColumn, IndpredColumn},
	}
}
//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgStatXactUserTables struct {
	Relid       *string
	Schemaname  *string
	Relname     *string
	SeqScan     *int64
	SeqTupRead  *int64
	IdxScan     *int64
	IdxTupFetch *int64
	NTupIns     *int64
	NTupUpd     *int64
	NTupDel     *int64
	NTupHotUpd  *int64
}

//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgSequence struct {
	Seqrelid     string
	Seqtypid     string
	Seqstart     int64
	Seqincrement int64
	Seqmax       int64
	Seqmin       int64
	Seqcache     int64
	Seqcycle     bool
}
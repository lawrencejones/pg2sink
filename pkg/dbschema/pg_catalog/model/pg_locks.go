//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgLocks struct {
	Locktype           *string
	Database           *string
	Relation           *string
	Page               *int32
	Tuple              *int16
	Virtualxid         *string
	Transactionid      *string
	Classid            *string
	Objid              *string
	Objsubid           *int16
	Virtualtransaction *string
	Pid                *int32
	Mode               *string
	Granted            *bool
	Fastpath           *bool
}

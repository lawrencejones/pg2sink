//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgStatProgressCreateIndex struct {
	Pid              *int32
	Datid            *string
	Datname          *string
	Relid            *string
	IndexRelid       *string
	Command          *string
	Phase            *string
	LockersTotal     *int64
	LockersDone      *int64
	CurrentLockerPid *int64
	BlocksTotal      *int64
	BlocksDone       *int64
	TuplesTotal      *int64
	TuplesDone       *int64
	PartitionsTotal  *int64
	PartitionsDone   *int64
}

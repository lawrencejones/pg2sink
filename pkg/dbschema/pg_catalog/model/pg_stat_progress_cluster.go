//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgStatProgressCluster struct {
	Pid               *int32
	Datid             *string
	Datname           *string
	Relid             *string
	Command           *string
	Phase             *string
	ClusterIndexRelid *string
	HeapTuplesScanned *int64
	HeapTuplesWritten *int64
	HeapBlksTotal     *int64
	HeapBlksScanned   *int64
	IndexRebuildCount *int64
}

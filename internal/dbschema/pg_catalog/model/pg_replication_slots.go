//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgReplicationSlots struct {
	SlotName          *string
	Plugin            *string
	SlotType          *string
	Datoid            *string
	Database          *string
	Temporary         *bool
	Active            *bool
	ActivePid         *int32
	Xmin              *string
	CatalogXmin       *string
	RestartLsn        *string
	ConfirmedFlushLsn *string
	WalStatus         *string
	SafeWalSize       *int64
}

//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type PgCursors struct {
	Name         *string
	Statement    *string
	IsHoldable   *bool
	IsBinary     *bool
	IsScrollable *bool
	CreationTime *time.Time
}

//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgIndex struct {
	Indexrelid     string
	Indrelid       string
	Indnatts       int16
	Indnkeyatts    int16
	Indisunique    bool
	Indisprimary   bool
	Indisexclusion bool
	Indimmediate   bool
	Indisclustered bool
	Indisvalid     bool
	Indcheckxmin   bool
	Indisready     bool
	Indislive      bool
	Indisreplident bool
	Indkey         string
	Indcollation   string
	Indclass       string
	Indoption      string
	Indexprs       *string
	Indpred        *string
}

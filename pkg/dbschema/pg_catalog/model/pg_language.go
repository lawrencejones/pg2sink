//
// Code generated by go-jet DO NOT EDIT.
// Generated at Tuesday, 12-May-20 07:59:32 BST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type PgLanguage struct {
	Oid           string
	Lanname       string
	Lanowner      string
	Lanispl       bool
	Lanpltrusted  bool
	Lanplcallfoid string
	Laninline     string
	Lanvalidator  string
	Lanacl        *string
}

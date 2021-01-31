// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Tables HTTP server types
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package server

import (
	tables "github.com/lawrencejones/pgsink/api/gen/tables"
)

// ListResponseBody is the type of the "Tables" service "List" endpoint HTTP
// response body.
type ListResponseBody []*TableResponse

// TableResponse is used to define fields on response body types.
type TableResponse struct {
	// Postgres table schema
	Schema string `form:"schema" json:"schema" xml:"schema"`
	// Postgres table name
	Name string `form:"name" json:"name" xml:"name"`
	// True if this table is already streaming
	Published bool `form:"published" json:"published" xml:"published"`
}

// NewListResponseBody builds the HTTP response body from the result of the
// "List" endpoint of the "Tables" service.
func NewListResponseBody(res []*tables.Table) ListResponseBody {
	body := make([]*TableResponse, len(res))
	for i, val := range res {
		body[i] = marshalTablesTableToTableResponse(val)
	}
	return body
}

// NewListPayload builds a Tables service List endpoint payload.
func NewListPayload(schema string) *tables.ListPayload {
	v := &tables.ListPayload{}
	v.Schema = schema

	return v
}

package tests

import "github.com/jcapobianco-cbi/easyjson"

//easyjson:json
type StructWithUnknownsProxy struct {
	easyjson.UnknownFieldsProxy

	Field1 string
}

//easyjson:json
type StructWithUnknownsProxyWithOmitempty struct {
	easyjson.UnknownFieldsProxy

	Field1 string `json:",omitempty"`
}

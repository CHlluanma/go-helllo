package common

import "encoding/json"

type clientRequest struct {
	Method string         `json:"method,omitempty"`
	Params [1]interface{} `json:"params,omitempty"`
	Id     uint64         `json:"id,omitempty"`
}

type serverRequest struct {
	Method string           `json:"method,omitempty"`
	Params *json.RawMessage `json:"params,omitempty"`
	Id     *json.RawMessage `json:"id,omitempty"`
}

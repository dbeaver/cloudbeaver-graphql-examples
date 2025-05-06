package graphql

import "encoding/json"

type Object map[string]any

// https://graphql.org/learn/serving-over-http/#post-request-and-body
type Request struct {
	Query     string `json:"query"`
	Variables Object `json:"variables,omitempty"`
}

// https://graphql.org/learn/serving-over-http/#response-format
type Response struct {
	Data   json.RawMessage
	Errors json.RawMessage
}

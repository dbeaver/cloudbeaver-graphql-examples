package graphql

// https://graphql.org/learn/serving-over-http/#post-request-and-body
type Request struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables,omitempty"`
}

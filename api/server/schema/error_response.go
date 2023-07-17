// Package schema contains request and response schemas of the server.
package schema

// Error contains fields of error response
type Error struct {
	Message string `json:"message"`
}

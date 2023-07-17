// Package models contains entities specific to
// the backend-test-golang application domain.
package models

import "math/big"

// Fact contains fields of fact.
type Fact struct {
	ID     string  `json:"id"`
	Text   string  `json:"text"`
	Number big.Int `json:"number"`
	Found  bool    `json:"found"`
	Type   string  `json:"type"`
}

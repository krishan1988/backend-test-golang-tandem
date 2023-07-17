// Package schema contains request and response schemas of the server.
package schema

import "github.com/KryptoKnight/backend-test-golang/models"

// RetrieveFactsResponse contains response data related to retrieve facts.
type RetrieveFactsResponse struct {
	Facts []models.Fact `json:"facts"`
}

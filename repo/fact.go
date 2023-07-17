// Package repo contains the interface layer for the database.
// In this package different databases can be implemented to interface them.
package repo

import (
	"context"

	"github.com/KryptoKnight/backend-test-golang/models"
)

// FactRepo in responsible for managing facts database transaction
type FactRepo interface {
	// GetAll returns list of facts by given filters.
	GetAll(ctx context.Context, filter models.FactFilter) ([]models.Fact, error)
}

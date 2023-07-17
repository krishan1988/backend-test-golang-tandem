// Package repo contains the interface layer for the database.
// In this package different databases can be implemented to interface them.
package repo

import (
	"context"
	"reflect"

	"github.com/KryptoKnight/backend-test-golang/models"
)

// NewFactMockRepo returns an instance of mock db implementation for FactRepo.
func NewFactMockRepo() *FactMockStore {
	return &FactMockStore{}
}

// FactMockStore a mock implementation for FactRepo.
type FactMockStore struct {
	getAllData struct {
		filter models.FactFilter
		facts  []models.Fact
		err    error
	}
}

// GetAll returns list of facts by given filters.
func (fms *FactMockStore) GetAll(ctx context.Context, filter models.FactFilter) ([]models.Fact, error) {
	if reflect.DeepEqual(fms.getAllData.filter, filter) {
		return fms.getAllData.facts, fms.getAllData.err
	}
	return nil, nil
}

// SetGetAll sets mock data for GetAll function.
func (fms *FactMockStore) SetGetAll(filter models.FactFilter, facts []models.Fact, err error) {
	fms.getAllData = struct {
		filter models.FactFilter
		facts  []models.Fact
		err    error
	}{filter: filter, facts: facts, err: err}
}

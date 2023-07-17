// Package service contains functionalities which are related to the business logics.
package service

import (
	"context"

	"github.com/KryptoKnight/backend-test-golang/models"
	"github.com/KryptoKnight/backend-test-golang/repo"
	"github.com/rs/zerolog"
)

// FactService manage all facts related business functionalities.
type FactService interface {

	// Retrieve is responsible for retrieving data from the database and do business functionalities.
	Retrieve(ctx context.Context, filter models.FactFilter) ([]models.Fact, error)
}

// factService implementation of facts.
type factService struct {
	logger zerolog.Logger
	repo   repo.FactRepo
}

// NewFactService creates an new instance of FactService.
func NewFactService(repo repo.FactRepo, logger zerolog.Logger) FactService {
	return &factService{
		repo:   repo,
		logger: logger,
	}
}

// Retrieve is responsible for retrieving data from the database and do business functionalities.
// Note: in this implementation, any valid business functionalities couldn't be find.
func (fs *factService) Retrieve(ctx context.Context, filter models.FactFilter) ([]models.Fact, error) {
	facts, err := fs.repo.GetAll(ctx, filter)
	if err != nil {
		fs.logger.Err(err).Msg("unable to get facts from database")
		return nil, err
	}
	return facts, nil
}

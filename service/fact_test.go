// Package service_test contains unit tests of service package.
package service_test

import (
	"context"
	"errors"
	"math/big"
	"os"
	"testing"

	"github.com/KryptoKnight/backend-test-golang/models"
	"github.com/KryptoKnight/backend-test-golang/repo"
	"github.com/KryptoKnight/backend-test-golang/service"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

// TestFactService_Retrieve test Retrieve function of the FactService.
func TestFactService_Retrieve(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		Desc          string
		Filter        models.FactFilter
		ExpectedFacts []models.Fact
		ExpectedErr   error
	}{
		{
			Desc: "Test when database return facts",
			Filter: models.FactFilter{
				Types: []string{"trivia"},
				Found: "true",
			},
			ExpectedFacts: []models.Fact{
				{
					Text:   "93 is that approximate distance in millions of miles the Sun is away from the Earth.",
					Number: *big.NewInt(0),
					Found:  true,
					Type:   "trivia",
				},
				{
					Text:   "137 is the California Penal Code for \"Offer bribe to influence testimony\".",
					Number: *big.NewInt(127),
					Found:  true,
					Type:   "trivia",
				},
			},
			ExpectedErr: nil,
		},
		{
			Desc: "Test when database return error",
			Filter: models.FactFilter{
				Types: []string{"trivia"},
				Found: "true",
			},
			ExpectedFacts: nil,
			ExpectedErr:   errors.New("sample return"),
		},
	}

	for _, test := range tests {
		writer := zerolog.New(os.Stdout)

		factsRepo := repo.NewFactMockRepo()
		factsRepo.SetGetAll(test.Filter, test.ExpectedFacts, test.ExpectedErr)

		actualFacts, actualErr := service.NewFactService(factsRepo, writer).Retrieve(ctx, test.Filter)

		assert.Equal(t, test.ExpectedFacts, actualFacts, test.Desc)
		assert.Equal(t, test.ExpectedErr, actualErr, test.Desc)
	}
}

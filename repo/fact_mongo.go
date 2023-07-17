// Package repo contains the interface layer for the database.
// In this package different databases can be implemented to interface them.
package repo

import (
	"context"
	"fmt"
	"math/big"

	"github.com/KryptoKnight/backend-test-golang/models"
	"github.com/KryptoKnight/backend-test-golang/mongo"
	"go.mongodb.org/mongo-driver/bson"
	driverMongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const factsCollectionName = "facts" // name of the collection of documents.

// NewFactMongoRepo returns an instance of mongo db implementation for FactRepo.
func NewFactMongoRepo() FactRepo {
	return &factMongoStore{
		collection: mongo.GetDB().Collection(factsCollectionName),
	}
}

// factMongoStore a mongo db implementation for FactRepo.
type factMongoStore struct {
	collection *driverMongo.Collection
}

// GetAll returns list of facts by given filters.
func (fms *factMongoStore) GetAll(ctx context.Context, filter models.FactFilter) ([]models.Fact, error) {
	facts := make([]models.Fact, 0)
	factDocs := make([]FactDoc, 0)

	m := bson.M{}

	if value, ok := filter.GetFound(); ok {
		m["found"] = value
	}

	if values, ok := filter.GetTypes(); ok {
		m["type"] = bson.M{
			"$in": values,
		}
	}

	limit := int64(filter.Limit)
	skip := int64(filter.Page*filter.Limit - filter.Limit)
	findOptions := options.FindOptions{Limit: &limit, Skip: &skip}

	cursor, err := fms.collection.Find(ctx, m, &findOptions)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &factDocs); err != nil {
		return facts, err
	}

	for _, doc := range factDocs {
		number := big.NewInt(0)
		number.SetString(fmt.Sprintf("%v", doc.Number), 10)

		facts = append(facts, models.Fact{
			ID:     doc.ID.Hex(),
			Text:   doc.Text,
			Type:   doc.Type,
			Number: *number,
			Found:  doc.Found,
		})
	}
	return facts, nil
}

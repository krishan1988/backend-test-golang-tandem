// Package repo contains the interface layer for the database.
// In this package different databases can be implemented to interface them.
package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

// FactDoc a model of the facts' document.
type FactDoc struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Text   string             `bson:"text,omitempty"`
	Number any                `bson:"number,omitempty"`
	Found  bool               `bson:"found,omitempty"`
	Type   string             `bson:"type,omitempty"`
}

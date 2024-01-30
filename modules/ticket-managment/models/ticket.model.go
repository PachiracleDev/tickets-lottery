package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ticket struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Code      string             `bson:"code"`
	Status    bool               `bson:"status"`
	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}

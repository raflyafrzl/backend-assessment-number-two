package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Psikolog struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}

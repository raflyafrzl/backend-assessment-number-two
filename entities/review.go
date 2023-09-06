package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Message string             `json:"message" bson:"message"`
	PsiId   primitive.ObjectID `json:"psikolog_id" bson:"psikolog_id"`
	Rating  int                `json:"rating" bson:"rating"`
}

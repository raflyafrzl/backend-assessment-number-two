package entities

type Review struct {
	Id      string `json:"_id" bson:"_id"`
	PsiId   string `json:"psikolog_id" bson:"psikolog_id"`
	Rating  int    `json:"rating" bson:"rating"`
	Message string `json:"message" bson:"message"`
}

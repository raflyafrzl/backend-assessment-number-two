package model

type CreateReviewModel struct {
	Rating  int    `json:"rating"`
	Message string `json:"message"`
	PsiId   string `json:"psikolog_id"`
}

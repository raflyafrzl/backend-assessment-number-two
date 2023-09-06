package entities

type PsikologReview struct {
	Id      string           `json:"_id"`
	Name    string           `json:"name"`
	Review  []map[string]any `json:"review"`
	Average float32          `json:"average_rating"`
}

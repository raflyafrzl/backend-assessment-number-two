package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"teduh-mongodb-assessment/contract"
	"teduh-mongodb-assessment/entities"
	"teduh-mongodb-assessment/model"

	"github.com/go-chi/chi/v5"
)

type PsikologController struct {
	service contract.PsikologUseCase
}

func NewPsikologController(s *contract.PsikologUseCase) *PsikologController {

	return &PsikologController{
		service: *s,
	}
}

func (p *PsikologController) Route(r chi.Router) {
	r.Get("/", p.List)
	r.Post("/", p.Create)
}

func (p *PsikologController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request entities.Psikolog

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 500,
			Error:      err.Error(),
		})
	}

	_ = json.Unmarshal(body, &request)

	var rawResponse map[string]any = map[string]any{
		"Status":     "Success",
		"StatusCode": 201,
		"Message":    "Data has been successfully created",
	}

	p.service.Create(request.Name)

	response, _ := json.Marshal(rawResponse)
	w.WriteHeader(201)

	w.Write(response)

}
func (p *PsikologController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []entities.PsikologReview = p.service.List()

	var rawResponse map[string]any = map[string]any{
		"result": results,
	}

	var response []byte

	response, _ = json.Marshal(rawResponse)

	w.Write(response)

}

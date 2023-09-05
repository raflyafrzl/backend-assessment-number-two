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
	service contract.PsikologService
}

func NewPsikologController(s *contract.PsikologService) *PsikologController {

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

	r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 500,
			Error:      err.Error(),
		})
	}

	_ = json.Unmarshal(body, &request)

	p.service.Create(request.Name)

	w.Write([]byte("Success"))

}
func (p *PsikologController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []entities.Psikolog = p.service.List()

	var rawResponse map[string]any = map[string]any{
		"Status":  "Success",
		"Message": "test",
		"result":  results,
	}

	var response []byte

	response, _ = json.Marshal(rawResponse)

	w.Write(response)

}

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

type ReviewController struct {
	ucase contract.ReviewUseCase
}

func (c *ReviewController) Route(r chi.Router) {

	r.Get("/", c.List)
	r.Post("/", c.Create)
}

func NewReviewController(ucase *contract.ReviewUseCase) *ReviewController {
	return &ReviewController{
		ucase: *ucase,
	}
}

func (c *ReviewController) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []entities.Review

	results = c.ucase.List()

	var response []byte
	response, _ = json.Marshal(results)

	w.Write(response)
}

func (c *ReviewController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request model.CreateReviewModel

	var body []byte
	var err error

	defer r.Body.Close()
	body, err = io.ReadAll(r.Body)

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: 400,
			Error:      "Invalid payload request",
		})
	}

	_ = json.Unmarshal(body, &request)

	c.ucase.Create(request)

	var rawResponse map[string]any = map[string]any{
		"Status":     "Success",
		"StatusCode": 201,
		"Message":    "You have been successfully created data",
	}

	response, _ := json.Marshal(rawResponse)

	w.Write(response)
}

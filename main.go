package main

import (
	"encoding/json"
	"net/http"
	"os"
	"teduh-mongodb-assessment/contract"
	"teduh-mongodb-assessment/database"
	"teduh-mongodb-assessment/delivery"
	"teduh-mongodb-assessment/middlewares"
	"teduh-mongodb-assessment/repository"
	"teduh-mongodb-assessment/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	var err error

	err = godotenv.Load()

	if err != nil {
		panic(err)
	}

	db := database.InitMongoDB()

	var r *chi.Mux = chi.NewRouter()

	r.Use(middlewares.RecoveryMiddleware)

	var psikologRepository contract.PsikologRepository = repository.NewPsikologRepostory(db)
	var psikologUseCase contract.PsikologUseCase = usecase.NewPsikologUseCase(&psikologRepository)
	var psikologController = delivery.NewPsikologController(&psikologUseCase)

	var reviewRepository contract.ReviewRepository = repository.NewReviewRepository(db)
	var reviewUseCase contract.ReviewUseCase = usecase.NewReviewUseCase(&reviewRepository)
	var reviewController = delivery.NewReviewController(&reviewUseCase)

	r.Route("/api/v1/review", reviewController.Route)
	r.Route("/api/v1/psikolog", psikologController.Route)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var rawResponse map[string]any = map[string]any{
			"Status": "Failed",
			"Error":  "Route not found",
		}

		var response []byte
		response, _ = json.Marshal(rawResponse)

		w.Write(response)
	})

	http.ListenAndServe(os.Getenv("PORT"), r)
}

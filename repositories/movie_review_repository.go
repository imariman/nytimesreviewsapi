package repositories

import (
	"encoding/json"
	"fmt"
	"log"
	"nytimesapi/models"
	"nytimesapi/utils"
	"os"
)

type MovieReviewRepository interface {
	SearchMovieReviews(*models.MovieReviewSearchParams) ([]models.MovieReview, error)
}

type movieReviewRepository struct {
}

func NewMovieReviewRepository() MovieReviewRepository {
	return &movieReviewRepository{}
}

func (repo *movieReviewRepository) SearchMovieReviews(movieReviewSearchParams *models.MovieReviewSearchParams) ([]models.MovieReview, error) {
	log.Println("Query: ", fmt.Sprint("https://api.nytimes.com/svc/movies/v2/reviews/search.json?", movieReviewSearchParams.GetQueryParams(), "api-key=", os.Getenv("NY_API_KEY")))
	_, body, err := utils.Get(fmt.Sprint("https://api.nytimes.com/svc/movies/v2/reviews/search.json?", movieReviewSearchParams.GetQueryParams(), "api-key=", os.Getenv("NY_API_KEY")))
	if err != nil {
		return nil, err
	}
	data := &models.MovieReviewHttpAnswer{}
	err = json.Unmarshal([]byte(string(body)), &data)
	if err != nil {
		return nil, err
	}
	return data.Results, nil
}

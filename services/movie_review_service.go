package services

import (
	"nytimesapi/models"
	"nytimesapi/repositories"
)

type MovieReviewService interface {
	SearchMovieReviews(*models.MovieReviewSearchParams) ([]models.MovieReview, error)
}

type movieReviewService struct {
	MovieReviewRepository repositories.MovieReviewRepository
}

func NewMovieReviewService(repo repositories.MovieReviewRepository) MovieReviewService {
	return &movieReviewService{MovieReviewRepository: repo}
}

func (service movieReviewService) SearchMovieReviews(movieReviewSearchParams *models.MovieReviewSearchParams) ([]models.MovieReview, error) {
	return service.MovieReviewRepository.SearchMovieReviews(movieReviewSearchParams)
}

package controllers

import (
	"log"
	"net/http"

	"nytimesapi/models"
	"nytimesapi/services"

	"github.com/gin-gonic/gin"
)

type MovieReviewController interface {
	SearchReviews(*gin.Context)
}

type movieReviewController struct {
	movieReviewService services.MovieReviewService
}

func NewMovieReviewController(movieReviewService services.MovieReviewService) MovieReviewController {
	return &movieReviewController{movieReviewService}
}

func (controller *movieReviewController) SearchReviews(c *gin.Context) {
	params := models.GetMovieReviewSearchParams(c)
	reviews, err := controller.movieReviewService.SearchMovieReviews(params)
	if err != nil {
		log.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, reviews)
}

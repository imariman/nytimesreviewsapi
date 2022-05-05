package main

import (
	"log"
	"nytimesapi/controllers"
	"nytimesapi/repositories"
	"nytimesapi/services"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Setup movie review repository
	movieReviewRepository := repositories.NewMovieReviewRepository()

	// Setup movie review service
	movieReviewService := services.NewMovieReviewService(movieReviewRepository)

	
	//	Setup movie review controller
	movieReviewController := controllers.NewMovieReviewController(movieReviewService)

	// Setup Gin Router, routes and logger and recovery middlewares
	router := gin.Default()
	router.GET("/api/v1/search", movieReviewController.SearchReviews)
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}

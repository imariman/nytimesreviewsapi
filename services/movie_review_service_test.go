package services

import (
	"nytimesapi/models"
	"nytimesapi/repositories"
	"os"
	"strings"
	"testing"
)

func init() {
	os.Setenv("NY_API_KEY", "MyEHyS6o5z0ZHAzw69vExZxdwY7Q6G6w")
}

func TestSearchMovieReviewService(t *testing.T) {
	movieReviewRepository := repositories.NewMovieReviewRepository()
	movieReviewService := NewMovieReviewService(movieReviewRepository)
	t.Run("GetReviewByMovieTitle", func(t *testing.T) {
		testTitle := "lebowski"
		reviews, err := movieReviewService.SearchMovieReviews(&models.MovieReviewSearchParams{
			MovieTitle: testTitle,
		})
		if err != nil {
			t.Error(err)
		}
		t.Log(len(reviews), " reviews fetched")
		for _, v := range reviews {
			if !strings.Contains(strings.ToLower(v.Headline), testTitle) {
				t.Error("GetReviewByMovieTitle FAILED. ", v.Headline, " does not contains ", testTitle)
				return
			}
		}
		t.Logf("GetByMovieTitle PASSED")
	})
	t.Run("GetReviewByReviewer", func(t *testing.T) {
		testReviewer := "jeannette catsoulis"
		reviews, err := movieReviewService.SearchMovieReviews(&models.MovieReviewSearchParams{
			ReviewerName: testReviewer,
		})
		if err != nil {
			t.Error(err)
		}
		t.Log(len(reviews), " reviews fetched")
		for _, v := range reviews {
			if !strings.Contains(strings.ToLower(v.Byline), testReviewer) {
				t.Error("GetByMovieTitle FAILED. ", v.Byline, " does not contains ", testReviewer)
				return
			}
		}
		t.Logf("GetByMovieTitle PASSED")
	})

}

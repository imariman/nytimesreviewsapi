package repositories

import (
	"nytimesapi/models"
	"os"
	"strings"
	"testing"
)

func init() {
	os.Setenv("NY_API_KEY", "MyEHyS6o5z0ZHAzw69vExZxdwY7Q6G6w")
}

func TestSearchMovieReviews(t *testing.T) {
	movieReviewRepository := NewMovieReviewRepository()
	t.Run("GetByMovieTitle", func(t *testing.T) {
		testTitle := "lebowski"
		reviews, err := movieReviewRepository.SearchMovieReviews(&models.MovieReviewSearchParams{
			MovieTitle: testTitle,
		})
		if err != nil {
			t.Error(err)
		}
		t.Log(len(reviews), " reviews fetched")
		for _, v := range reviews {
			if !strings.Contains(strings.ToLower(v.Headline), strings.ToLower(testTitle)) {
				t.Error(v.Headline, " does not contains ", testTitle)
			}
		}
	})
	t.Run("GetByReviewer", func(t *testing.T) {
		testReviewer := "jeannette catsoulis"
		reviews, err := movieReviewRepository.SearchMovieReviews(&models.MovieReviewSearchParams{
			ReviewerName: testReviewer,
		})
		if err != nil {
			t.Error(err)
		}
		t.Log(len(reviews), " reviews fetched")
		for _, v := range reviews {
			if !strings.Contains(strings.ToLower(v.Byline), strings.ToLower(testReviewer)) {
				t.Error(v.Byline, " does not contains ", testReviewer)
				break
			}
		}
	})
}

package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"nytimesapi/models"
	"nytimesapi/repositories"
	"nytimesapi/services"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("NY_API_KEY", "MyEHyS6o5z0ZHAzw69vExZxdwY7Q6G6w")
}

func TestMovieReviewController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	movieReviewRepository := repositories.NewMovieReviewRepository()
	movieReviewService := services.NewMovieReviewService(movieReviewRepository)
	movieReviewController := NewMovieReviewController(movieReviewService)
	testDateTime := "2019-12-31"
	t.Run("GetByTitle", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		testQuery := "lebowski"
		req, _ := http.NewRequest("GET", "/api/v1/search?title="+testQuery, nil)
		c.Request = req
		movieReviewController.SearchReviews(c)
		assert.Equal(t, http.StatusOK, w.Code)
		resBody := make([]models.MovieReview, 0)
		err := json.NewDecoder(w.Body).Decode(&resBody)
		assert.Nil(t, err)
		assert.Greater(t, len(resBody), 0)
		for _, v := range resBody {
			assert.Contains(t, strings.ToLower(v.DisplayTitle), strings.ToLower(testQuery))
			assert.LessOrEqual(t, v.PublicationDate, testDateTime)
		}
	})
	t.Run("GetByReviewer", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		testQuery := "Elisabeth Vincentelli"
		req, _ := http.NewRequest("GET", "/api/v1/search?reviewer="+testQuery, nil)
		c.Request = req
		movieReviewController.SearchReviews(c)
		assert.Equal(t, http.StatusOK, w.Code)
		resBody := make([]models.MovieReview, 0)
		err := json.NewDecoder(w.Body).Decode(&resBody)
		assert.Nil(t, err)
		assert.Greater(t, len(resBody), 0)
		for _, v := range resBody {
			assert.Equal(t, strings.ToLower(v.Byline), strings.ToLower(testQuery))
			assert.LessOrEqual(t, v.PublicationDate, testDateTime)
		}
	})
	t.Run("GetByPublicationDate", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		testDateRange := "2022-05-03:2022-05-04"
		req, _ := http.NewRequest("GET", "/api/v1/search?publication-date="+testDateRange, nil)
		c.Request = req
		movieReviewController.SearchReviews(c)
		assert.Equal(t, http.StatusOK, w.Code)
		resBody := make([]models.MovieReview, 0)
		err := json.NewDecoder(w.Body).Decode(&resBody)
		assert.Nil(t, err)
		assert.Greater(t, len(resBody), 0)
		if err != nil {
			t.Error("invalid datetime: ", err)
			return
		}
		arr := strings.Split(testDateRange, ":")
		for _, v := range resBody {
			assert.GreaterOrEqual(t, v.PublicationDate, arr[0])
			assert.LessOrEqual(t, v.PublicationDate, arr[1])
		}
	})
	t.Run("GetByReviewerAndTitle", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		testReviewer := "Glenn Kenny"
		testTitle := "annual"
		req, _ := http.NewRequest("GET", "/api/v1/search?reviewer="+testReviewer+"&title="+testTitle, nil)
		c.Request = req
		movieReviewController.SearchReviews(c)
		assert.Equal(t, http.StatusOK, w.Code)
		resBody := make([]models.MovieReview, 0)
		err := json.NewDecoder(w.Body).Decode(&resBody)
		assert.Nil(t, err)
		assert.Greater(t, len(resBody), 0)
		for _, v := range resBody {
			assert.Contains(t, strings.ToLower(v.DisplayTitle), strings.ToLower(testTitle))
			assert.Equal(t, strings.ToLower(v.Byline), strings.ToLower(testReviewer))
			assert.LessOrEqual(t, v.PublicationDate, testDateTime)
		}
	})
	t.Run("GetByReviewerAndTitleNotFound", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		testReviewer := "Beatrice Loayza"
		testTitle := "teo"
		req, _ := http.NewRequest("GET", "/api/v1/search?reviewer="+testReviewer+"&title="+testTitle, nil)
		c.Request = req
		movieReviewController.SearchReviews(c)
		assert.Equal(t, http.StatusOK, w.Code)
		resBody := make([]models.MovieReview, 0)
		err := json.NewDecoder(w.Body).Decode(&resBody)
		assert.Nil(t, err)
		assert.Equal(t, len(resBody), 0)
	})
}

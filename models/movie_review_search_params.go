package models

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type MovieReviewSearchParams struct {
	ReviewerName    string
	MovieTitle      string
	PublicationDate string
}

func GetMovieReviewSearchParams(c *gin.Context) *MovieReviewSearchParams {
	return &MovieReviewSearchParams{
		ReviewerName:    c.Query("reviewer"),
		MovieTitle:      c.Query("title"),
		PublicationDate: c.Query("publication-date"),
	}
}

func (params MovieReviewSearchParams) GetQueryParams() string {
	var sb strings.Builder
	if params.MovieTitle != "" {
		sb.Write([]byte("query="))
		sb.Write([]byte(params.MovieTitle))
	}
	if params.PublicationDate != "" {
		if sb.Len() > 0 {
			sb.Write([]byte("&"))
		}
		sb.Write([]byte("publication-date="))
		sb.Write([]byte(params.PublicationDate))
	} else {
		if sb.Len() > 0 {
			sb.Write([]byte("&"))
		}
		sb.Write([]byte("publication-date="))
		sb.Write([]byte("1900-01-01:2019-12-31"))
	}
	if params.ReviewerName != "" {
		if sb.Len() > 0 {
			sb.Write([]byte("&"))
		}
		sb.Write([]byte("reviewer="))
		sb.Write([]byte(params.ReviewerName))
	}
	if sb.Len() > 0 {
		sb.Write([]byte("&"))
	}
	return sb.String()
}

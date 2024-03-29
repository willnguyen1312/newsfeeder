package handlers

import (
	"net/http"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// NewsfeedPost handle get request to ping route
func NewsfeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}
		c.Bind(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}

		feed.Add(item)

		c.Status(http.StatusNoContent)
	}
}

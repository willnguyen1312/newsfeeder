package main

import (
	"newsfeeder/httpd/handlers"
	"newsfeeder/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handlers.PingGet)
	r.GET("/newsfeed", handlers.NewsfeedGet(feed))
	r.POST("/newsfeed", handlers.NewsfeedPost(feed))

	r.Run()
}

package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesusbibieca/go-url-shortener/handler"
	"github.com/jesusbibieca/go-url-shortener/store"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// Note that store initialization happens here
	store.InitializeStore()

	err := router.Run(":8080")

	if err != nil {
		panic(fmt.Sprintf("Failed to start server %v", err))
	}
}

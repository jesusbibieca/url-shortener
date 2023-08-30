package main

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesusbibieca/url-shortener/environment"
	"github.com/jesusbibieca/url-shortener/handler"
	"github.com/jesusbibieca/url-shortener/store"
)


func main() {
	config, err := environment.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Error loading config: %v", err))
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Service is up and running 🚀"})
	})

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err = router.Run(config.AppAddress)

	if err != nil {
		panic(fmt.Sprintf("Failed to start server %v", err))
	}
}

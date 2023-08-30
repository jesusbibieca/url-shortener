package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesusbibieca/url-shortener/shortener"
	"github.com/jesusbibieca/url-shortener/store"
)

type ShortUrlCreateRequest struct {
	Url    string `json:"url"`
	UserId string `json:"userId"`
}

const BASE_URL = "http://localhost:8080/"

func CreateShortUrl(c *gin.Context) {
	var shortUrlRequest ShortUrlCreateRequest
	if err := c.ShouldBindJSON(&shortUrlRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	shortUrl := shortener.GenerateShortLink(shortUrlRequest.Url, shortUrlRequest.UserId)
	store.SaveUrlMapping(shortUrl, shortUrlRequest.Url, shortUrlRequest.UserId)

	c.JSON(http.StatusOK, gin.H{
		"shortUrl": BASE_URL + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	url := store.RetrieveInitialUrl(shortUrl)

	c.Redirect(http.StatusTemporaryRedirect, url)
}

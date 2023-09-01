package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/", server.ping)
	router.GET("/:shortUrl", server.getShortUrl)

	router.POST("/shorten", server.createShortUrl)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	// Consider checking the error type using a switch statement
	// and return a different status code for different types of errors.
	return gin.H{"error": err.Error()}
}

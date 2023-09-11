package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/", server.ping)

	// Url routes
	router.GET("/r", server.getPagedUrls)
	router.GET("/r/:shortUrl", server.getShortUrl)
	router.POST("/shorten", server.createShortUrl)
	router.PATCH("/r/:shortUrl", server.updateShortUrl)
	router.DELETE("/r/:shortUrl", server.deleteShortUrl)

	// User routes
	router.GET("/users", server.getPagedUsers)
	router.GET("/user/:id", server.getUser)
	router.POST("/users", server.createUser)
	router.DELETE("/user/:id", server.deleteUser)

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

package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jesusbibieca/url-shortener/authentication"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
	"github.com/jesusbibieca/url-shortener/environment"
)

type Server struct {
	authTokenMaker authentication.PasetoMaker
	config         environment.Configuration
	router         *gin.Engine
	store          *db.Store
}

func NewServer(store *db.Store, config environment.Configuration) (*Server, error) {
	authTokenMaker, err := authentication.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create authTokenMaker %w", err)
	}

	server := &Server{
		store:          store,
		authTokenMaker: *authTokenMaker,
	}
	router := gin.Default()

	// Health check
	router.GET("/", server.ping)

	// Authentication
	router.POST("/auth/login", server.login)
	router.POST("/auth/register", server.createUser)

	authenticatedRoutes := router.Group("/").Use(authMiddleware(server.authTokenMaker))

	// Url routes
	authenticatedRoutes.GET("/r", server.getPagedUrls)
	authenticatedRoutes.POST("/r", server.createShortUrl)
	authenticatedRoutes.GET("/r/:shortUrl", server.getShortUrl)
	authenticatedRoutes.PATCH("/r/:shortUrl", server.updateShortUrl)
	authenticatedRoutes.DELETE("/r/:shortUrl", server.deleteShortUrl)

	// User routes
	authenticatedRoutes.GET("/users", server.getPagedUsers)
	authenticatedRoutes.GET("/user/:id", server.getUser)
	authenticatedRoutes.DELETE("/user/:id", server.deleteUser)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	// Consider checking the error type using a switch statement
	// and return a different status code for different types of errors.
	return gin.H{"error": err.Error()}
}

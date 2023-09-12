package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/jesusbibieca/url-shortener/db/sqlc"
	"github.com/jesusbibieca/url-shortener/environment"
	"github.com/jesusbibieca/url-shortener/helpers"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store *db.Store) *Server {
	config := environment.Configuration{
		TokenSymmetricKey:   helpers.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(store, config)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

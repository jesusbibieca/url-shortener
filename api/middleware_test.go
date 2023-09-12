package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jesusbibieca/url-shortener/authentication"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker authentication.PasetoMaker,
	authorizationType string,
	userID int32,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(userID, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	userID := int32(10)

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, authTokenMaker authentication.PasetoMaker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, authTokenMaker authentication.PasetoMaker) {
				addAuthorization(t, request, authTokenMaker, authorizationType, userID, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, authTokenMaker authentication.PasetoMaker) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, authTokenMaker authentication.PasetoMaker) {
				addAuthorization(t, request, authTokenMaker, "unsupported", userID, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InvalidAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, authTokenMaker authentication.PasetoMaker) {
				addAuthorization(t, request, authTokenMaker, "", userID, time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "ExpiredAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, authTokenMaker authentication.PasetoMaker) {
				addAuthorization(t, request, authTokenMaker, authorizationHeaderKey, userID, -time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			server := newTestServer(t, nil)
			// server.router.Use(authMiddleware(server.authTokenMaker))
			server.router.GET("/test/auth", authMiddleware(server.authTokenMaker), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{})
			})

			request, err := http.NewRequest(http.MethodGet, "/test/auth", nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.authTokenMaker)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

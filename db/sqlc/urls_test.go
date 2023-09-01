package db

import (
	"context"
	"testing"

	"github.com/jesusbibieca/url-shortener/helpers"
	"github.com/stretchr/testify/require"
)

func createNewShortUrl(t *testing.T) Url {
	user := createNewUser(t)

	args := CreateShortUrlParams{
		OriginalUrl: helpers.RandomUrl(),
		UserID:      user.ID,
	}

	shortUrl, err := testStore.CreateShortUrl(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, shortUrl)

	require.Equal(t, args.OriginalUrl, shortUrl.OriginalUrl)
	require.Equal(t, args.UserID, shortUrl.UserID)

	require.NotZero(t, shortUrl.ID)
	require.NotZero(t, shortUrl.CreatedAt)

	return shortUrl
}

func TestCreateShortUrl(t *testing.T) {
	createNewShortUrl(t)
}

func TestGetShortUrlByID(t *testing.T) {
	inputUrl := createNewShortUrl(t)
	dbUrl, err := testStore.GetShortUrlByID(context.Background(), inputUrl.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dbUrl)

	require.Equal(t, inputUrl.ID, dbUrl.ID)
	require.Equal(t, inputUrl.OriginalUrl, dbUrl.OriginalUrl)
	require.Equal(t, inputUrl.ShortUrl, dbUrl.ShortUrl)
	require.Equal(t, inputUrl.UserID, dbUrl.UserID)
	require.WithinDuration(t, inputUrl.CreatedAt.Time, dbUrl.CreatedAt.Time, 0)
}

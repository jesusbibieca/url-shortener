package authentication

import (
	"testing"
	"time"

	"github.com/jesusbibieca/url-shortener/helpers"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(helpers.RandomString(32))
	require.NoError(t, err)

	userID := int32(helpers.RandomInt(1, 10))
	duration := time.Minute

	issuedAt := time.Now()
	expireAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(userID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, userID, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expireAt, payload.ExpireAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	maker, err := NewPasetoMaker(helpers.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(int32(10), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, "token has expired")
	require.Nil(t, payload)
}

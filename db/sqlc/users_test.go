package db

import (
	"context"
	"testing"

	"github.com/jesusbibieca/url-shortener/helpers"
	"github.com/stretchr/testify/require"
)

func createNewUser(t *testing.T) User {
	hashedPassword, err := helpers.HashPassword("password")
	require.NoError(t, err)

	args := CreateUserParams{
		Username: helpers.RandomString(6),
		Email:    helpers.RandomEmail(),
		Password: hashedPassword,
	}

	user, err := testStore.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.Password, user.Password)
	require.Equal(t, args.Email, user.Email)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createNewUser(t)
}

func TestGetUser(t *testing.T) {
	inputUser := createNewUser(t)
	dbUser, err := testStore.GetUser(context.Background(), inputUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, dbUser)

	require.Equal(t, inputUser.Username, dbUser.Username)
	require.Equal(t, inputUser.Password, dbUser.Password)
	require.Equal(t, inputUser.Email, dbUser.Email)
	require.WithinDuration(t, inputUser.CreatedAt.Time, dbUser.CreatedAt.Time, 0)
}

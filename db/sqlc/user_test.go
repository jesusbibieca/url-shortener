package db

import (
	"context"
	"testing"

	"github.com/jesusbibieca/url-shortener/helpers"
	"github.com/stretchr/testify/require"
)

func createNewUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       helpers.RandomString(6),
		HashedPassword: "hashed password",
		FullName:       helpers.RandomString(6) + " " + helpers.RandomString(6),
		Email:          helpers.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createNewUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createNewUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, 0)
}

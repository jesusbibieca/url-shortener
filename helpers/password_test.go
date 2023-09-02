package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = VerifyPassword(password, hashedPassword)
	require.NoError(t, err)

	err = VerifyPassword("wrong password", hashedPassword)
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestHashSameString(t *testing.T) {
	password := RandomString(8)

	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	// Should be different
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}

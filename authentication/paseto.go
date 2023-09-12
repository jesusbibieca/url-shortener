package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

var (
	ErrInvalidToken = errors.New("provided token is not valid")
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

type Payload struct {
	UserID   int32     `json:"userId"`
	IssuedAt time.Time `json:"issueAt"`
	ExpireAt time.Time `json:"expireAt"`
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpireAt) {
		return errors.New("token has expired")
	}
	return nil
}

func NewPasetoMaker(symmetricKey string) (*PasetoMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size must be %d: ", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (m *PasetoMaker) CreateToken(userID int32, duration time.Duration) (string, error) {
	payload := &Payload{
		UserID:   userID,
		IssuedAt: time.Now(),
		ExpireAt: time.Now().Add(duration),
	}
	// Encrypt takes the key, the payload and the footer which is optional
	return m.paseto.Encrypt(m.symmetricKey, payload, nil)
}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := m.paseto.Decrypt(token, m.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()

	if err != nil {
		return nil, err
	}
	return payload, nil
}

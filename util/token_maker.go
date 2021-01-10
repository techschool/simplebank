package util

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

// TokenMaker manages token creation and validation
type TokenMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// NewTokenMaker returns a new TokenMaker
func NewTokenMaker(symmetricKey string) (*TokenMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be %d characters", chacha20poly1305.KeySize)
	}

	tokenMaker := &TokenMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return tokenMaker, nil
}

// Token contains the payload data of the token
type Token struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// CreateToken creates a new token for a username with a fixed valid duration
func (tokenMaker *TokenMaker) CreateToken(username string, duration time.Duration) (string, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	token := Token{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}

	return tokenMaker.paseto.Encrypt(tokenMaker.symmetricKey, token, nil)
}

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = fmt.Errorf("token is invalid")
	ErrExpiredToken = fmt.Errorf("token has expired")
)

// VerifyToken checks if the token is valid or not
func (tokenMaker *TokenMaker) VerifyToken(tokenString string) (*Token, error) {
	token := &Token{}

	err := tokenMaker.paseto.Decrypt(tokenString, tokenMaker.symmetricKey, token, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	if time.Now().After(token.ExpiredAt) {
		return nil, ErrExpiredToken
	}

	return token, nil
}

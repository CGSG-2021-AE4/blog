package api

import (
	"context"
	"io"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

type TokenClaims struct {
	Issuer         string // "iss"
	ExpirationTime int64  // "exp"
}
type Token string

type UserService interface {
	Login(ctx context.Context, username, password string) (Token, error)
	ValidateToken(ctx context.Context, token Token) (TokenClaims, error)
	Register(ctx context.Context, u *User) error
	io.Closer
}

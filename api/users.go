package api

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/google/uuid"
)

type TokenClaims struct {
	Issuer         string // "iss"
	ExpirationTime int64  // "exp"
}
type Token string

type UserService interface {
	Login(ctx context.Context, username, password string) (Token, error)
	Register(ctx context.Context, u *types.User) error
	Delete(ctx context.Context, id uuid.UUID) error

	ValidateToken(ctx context.Context, token Token) (TokenClaims, error)

	GetUser(ctx context.Context, id uuid.UUID) (*types.User, error)
	GetUserByName(ctx context.Context, username string) (*types.User, error)
	io.Closer
}

package api

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/google/uuid"
)

type TokenClaims struct {
	Issuer  uuid.UUID // "iss" user id
	ExpTime int64     // "exp"
}
type Token string

type UserPublic struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type UserService interface {
	Login(ctx context.Context, username, password string) (uuid.UUID, Token, error)
	Register(ctx context.Context, u *types.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, shift, limit uint) ([]UserPublic, error)

	ValidateToken(ctx context.Context, token Token) (TokenClaims, error)

	GetUser(ctx context.Context, id uuid.UUID) (*types.User, error)
	GetUserByName(ctx context.Context, username string) (*types.User, error)
	io.Closer
}

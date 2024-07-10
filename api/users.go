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

type UserStoreReader interface {
	GetUser(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByName(ctx context.Context, username string) (*User, error)
	DoExist(ctx context.Context, username string) (bool, error) // Is needed for registration
}

type UserStoreWriter interface {
	CreateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, Id uuid.UUID) error
}

type UserStore interface {
	UserStoreReader
	UserStoreWriter
	io.Closer
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

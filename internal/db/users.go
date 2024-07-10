package db

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/api"

	"github.com/google/uuid"
)

// Some common errors
const (
	ErrUserNotExists     = api.Error("user not exists")
	ErrUserAlreadyExists = api.Error("user already exists")
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

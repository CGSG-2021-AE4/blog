package db

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/internal/types"

	"github.com/google/uuid"
)

// Some common errors
const (
	ErrUserNotExists     = api.Error("user not exists")
	ErrUserAlreadyExists = api.Error("user already exists")
)

type UserStoreReader interface {
	GetUser(ctx context.Context, id uuid.UUID) (*types.User, error)
	GetUserByName(ctx context.Context, username string) (*types.User, error)
	DoExist(ctx context.Context, username string) (bool, error) // Is needed for registration
}

type UserStoreWriter interface {
	CreateUser(ctx context.Context, user *types.User) error
	DeleteUser(ctx context.Context, Id uuid.UUID) error
}

type UserStore interface {
	UserStoreReader
	UserStoreWriter
	io.Closer
}

package db

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/internal/types"

	"github.com/google/uuid"
)

const (
	ErrArticleNotFound      = api.Error("article not found")
	ErrArticleAlreadyExists = api.Error("article already exists")
)

type ArticleStore interface {
	List(ctx context.Context, limit int) ([]types.Article, error) // TODO add shift

	Get(ctx context.Context, id uuid.UUID) (types.Article, error)
	Create(ctx context.Context, a types.Article) error
	Update(ctx context.Context, a types.Article) error
	Delete(ctx context.Context, id uuid.UUID) error

	io.Closer
}

type ContentStore interface {
	Get(ctx context.Context, id uuid.UUID) ([]byte, error)
	Create(ctx context.Context, id uuid.UUID, content []byte) error
	Update(ctx context.Context, id uuid.UUID, content []byte) error
	Delete(ctx context.Context, id uuid.UUID) error

	io.Closer
}

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
	ListHeaders(ctx context.Context, limit int) ([]types.ArticleHeader, error) // TODO add shift
	GetArticle(ctx context.Context, Id uuid.UUID) (*types.Article, error)
	CreateArticle(ctx context.Context, a *types.Article) error
	DeleteArticle(ctx context.Context, Id uuid.UUID) error
	io.Closer
}

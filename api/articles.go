package api

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/google/uuid"
)

type ArticlesService interface {
	ListArticles(ctx context.Context, limit int) ([]types.ArticleHeader, error)
	GetArticle(ctx context.Context, id uuid.UUID) (*types.Article, error)
	CreateArticle(ctx context.Context, a *types.Article) error
	DeleteArticle(ctx context.Context, id uuid.UUID) error
	EditArticle(ctx context.Context, a *types.Article) error

	io.Closer
}

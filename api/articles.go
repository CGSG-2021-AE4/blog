package api

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/google/uuid"
)

type ArticlesService interface {
	ListArticles(ctx context.Context, limit int) ([]types.Article, error)
	GetArticle(ctx context.Context, id uuid.UUID) (types.Article, error)
	GetContent(ctx context.Context, id uuid.UUID) ([]byte, error)

	CreateArticle(ctx context.Context, descr types.ArticleDescr) (uuid.UUID, error)
	EditArticle(ctx context.Context, id uuid.UUID, descr types.ArticleDescr) error
	EditContent(ctx context.Context, id uuid.UUID, content []byte) error
	DeleteArticle(ctx context.Context, id uuid.UUID) error

	io.Closer
}

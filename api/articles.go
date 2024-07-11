package api

import (
	"context"
	"io"

	"github.com/CGSG-2021-AE4/blog/internal/types"
)

type ArticlesService interface {
	ListArticles(ctx context.Context, limit int) ([]types.ArticleHeader, error)
	io.Closer
}

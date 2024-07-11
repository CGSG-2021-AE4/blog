package app

import (
	"context"

	"github.com/CGSG-2021-AE4/blog/internal/db"
	"github.com/CGSG-2021-AE4/blog/internal/types"
)

type ArticlesService struct {
	artsStore db.ArticleStore
}

func NewArticlesService(artsStore db.ArticleStore) *ArticlesService {
	return &ArticlesService{
		artsStore: artsStore,
	}
}

func (svc *ArticlesService) ListArticles(ctx context.Context, limit int) ([]types.ArticleHeader, error) {
	return svc.artsStore.ListHeaders(ctx, limit)
}

func (svc *ArticlesService) Close() error {
	return nil
}

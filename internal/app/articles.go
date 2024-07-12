package app

import (
	"context"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/internal/db"
	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/google/uuid"
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

func (svc *ArticlesService) GetArticle(ctx context.Context, id uuid.UUID) (*types.Article, error) {
	return svc.artsStore.GetArticle(ctx, id)
}

func (svc *ArticlesService) CreateArticle(ctx context.Context, a *types.Article) error {
	return svc.artsStore.CreateArticle(ctx, a)
}

func (svc *ArticlesService) DeleteArticle(ctx context.Context, id uuid.UUID) error {
	return svc.artsStore.DeleteArticle(ctx, id)
}

func (svc *ArticlesService) EditArticle(ctx context.Context, a *types.Article) error {
	return api.ErrNotImplementedYet
}

func (svc *ArticlesService) Close() error {
	return nil
}

package app

import (
	"context"
	"fmt"

	"github.com/CGSG-2021-AE4/blog/internal/db"
	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/google/uuid"
)

type ArticlesService struct {
	artsStore    db.ArticleStore
	contentStore db.ContentStore
}

func NewArticlesService(artsStore db.ArticleStore, contentStore db.ContentStore) *ArticlesService {
	return &ArticlesService{
		artsStore:    artsStore,
		contentStore: contentStore,
	}
}

func (svc *ArticlesService) ListArticles(ctx context.Context, limit int) ([]types.Article, error) {
	return svc.artsStore.List(ctx, limit)
}

func (svc *ArticlesService) GetArticle(ctx context.Context, id uuid.UUID) (types.Article, error) {
	return svc.artsStore.Get(ctx, id)
}

func (svc *ArticlesService) CreateArticle(ctx context.Context, descr types.ArticleDescr) (uuid.UUID, error) {
	a := types.Article{
		ArticleDescr: descr,
		Id:           uuid.New(),
		ContentId:    uuid.New(),
	}
	if err := svc.artsStore.Create(ctx, a); err != nil {
		return uuid.Nil, err
	}
	if err := svc.contentStore.Create(ctx, a.ContentId, []byte{}); err != nil {
		if dErr := svc.artsStore.Delete(ctx, a.Id); dErr != nil {
			// What am I supposed to do in this situation
			return uuid.Nil, fmt.Errorf("error '%w' while error '%w'", dErr, err)
		}
		return uuid.Nil, err
	}
	return a.Id, nil
}

func (svc *ArticlesService) EditArticle(ctx context.Context, id uuid.UUID, descr types.ArticleDescr) error {
	a, err := svc.artsStore.Get(ctx, id)
	if err != nil {
		return fmt.Errorf("get article: %w", err)
	}
	a.ArticleDescr = descr
	if err := svc.artsStore.Update(ctx, a); err != nil {
		return fmt.Errorf("update article: %w", err)
	}
	return nil
}

func (svc *ArticlesService) EditContent(ctx context.Context, id uuid.UUID, content []byte) error {
	return svc.contentStore.Update(ctx, id, content)
}

func (svc *ArticlesService) DeleteArticle(ctx context.Context, id uuid.UUID) error {
	return svc.artsStore.Delete(ctx, id)
}

func (svc *ArticlesService) GetContent(ctx context.Context, id uuid.UUID) ([]byte, error) {
	return svc.contentStore.Get(ctx, id)
}

func (svc *ArticlesService) Close() error {
	return nil
}

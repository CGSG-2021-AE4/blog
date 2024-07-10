package json

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/internal/db"

	"github.com/google/uuid"
)

type ArticleStore struct {
	filename string

	mutex    sync.Mutex
	articles map[uuid.UUID]*db.Article
}

func NewArticleStore(filename string) (*ArticleStore, error) {
	as := ArticleStore{
		filename: filename,
		articles: make(map[uuid.UUID]*db.Article),
	}
	if err := as.load(); err != nil {
		return nil, err
	}
	return &as, nil
}

func (as *ArticleStore) load() error {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	bytes, err := os.ReadFile(as.filename)
	if err != nil {
		return err
	}
	articles := []db.ArticleJson{}
	if err := json.Unmarshal(bytes, &articles); err != nil {
		return nil
	}
	for _, u := range articles { // Check if this is memory friendly
		as.articles[u.Id] = &db.Article{
			Header:  u.ArticleHeader,
			Content: &u.ArticleContent,
		}
	}
	return nil
}

func (as *ArticleStore) save() error {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	articles := []db.ArticleJson{}
	for _, a := range as.articles {
		articles = append(articles, db.ArticleJson{ArticleHeader: a.Header, ArticleContent: *a.Content})
	}
	bytes, err := json.Marshal(articles)
	if err != nil {
		return err
	}
	return os.WriteFile(as.filename, bytes, 0777)
}

func (as *ArticleStore) ListHeaders(ctx context.Context, limit int) ([]db.ArticleHeader, error) {
	return nil, api.ErrNotImplementedYet
}
func (as *ArticleStore) GetArticle(ctx context.Context, Id uuid.UUID) (*db.Article, error) {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	if a := as.articles[Id]; a != nil {
		return a, nil
	}
	return nil, db.ErrArticleNotFound
}

func (as *ArticleStore) CreateArticle(ctx context.Context, a *db.Article) error {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	// For better safe
	if as.articles[a.Header.Id] != nil {
		return db.ErrArticleAlreadyExists
	}
	as.articles[a.Header.Id] = a
	return nil
}
func (as *ArticleStore) DeleteArticle(ctx context.Context, Id uuid.UUID) error {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	// For better safe
	if as.articles[Id] == nil {
		return db.ErrArticleNotFound
	}
	delete(as.articles, Id)
	return nil
}
func (as *ArticleStore) Close() error {
	return as.save()
}

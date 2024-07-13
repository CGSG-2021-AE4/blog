package json

import (
	"context"
	"encoding/json"
	"os"
	"sync"

	"github.com/CGSG-2021-AE4/blog/internal/db"
	"github.com/CGSG-2021-AE4/blog/internal/types"

	"github.com/google/uuid"
)

type ArticleStore struct {
	filename string

	mutex    sync.Mutex
	articles map[uuid.UUID]*types.Article
}

func NewArticleStore(filename string) (*ArticleStore, error) {
	as := ArticleStore{
		filename: filename,
		articles: make(map[uuid.UUID]*types.Article),
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
	articles := []types.Article{}
	if err := json.Unmarshal(bytes, &articles); err != nil {
		return nil
	}
	for _, a := range articles { // Check if this is memory friendly
		as.articles[a.Id] = &a
	}
	return nil
}

func (as *ArticleStore) save() error {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	articles := []types.Article{}
	for _, a := range as.articles {
		articles = append(articles, *a)
	}
	bytes, err := json.Marshal(articles)
	if err != nil {
		return err
	}
	return os.WriteFile(as.filename, bytes, 0777)
}

func (as *ArticleStore) List(ctx context.Context, limit int) ([]types.Article, error) {
	headers := []types.Article{}

	for _, a := range as.articles { //
		headers = append(headers, *a)
	}
	return headers, nil
}

func (as *ArticleStore) Get(ctx context.Context, Id uuid.UUID) (types.Article, error) {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	if a := as.articles[Id]; a != nil {
		return *a, nil
	}
	return types.NilArticle, db.ErrArticleNotFound
}

func (as *ArticleStore) Update(ctx context.Context, a types.Article) error {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	if found := as.articles[a.Id]; found != nil {
		found = &a
		return nil
	}
	return db.ErrArticleNotFound
}

func (as *ArticleStore) Create(ctx context.Context, a types.Article) error {
	as.mutex.Lock()
	defer as.mutex.Unlock()

	// For better safe
	if as.articles[a.Id] != nil {
		return db.ErrArticleAlreadyExists
	}
	as.articles[a.Id] = &a
	return nil
}

func (as *ArticleStore) Delete(ctx context.Context, Id uuid.UUID) error {
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

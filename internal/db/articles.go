package db

import (
	"context"
	"io"
	"webapp/api"

	"github.com/google/uuid"
)

const (
	ErrArticleNotFound      = api.Error("article not found")
	ErrArticleAlreadyExists = api.Error("article already exists")
)

type ArticleContent struct {
	Text string `json:"text"`
}

type ArticleHeader struct {
	Id    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

type Article struct {
	Header  ArticleHeader
	Content *ArticleContent // Pointer it is supposed to be long
}

type ArticleJson struct {
	ArticleHeader
	ArticleContent
}

type ArticleStore interface {
	ListHeaders(ctx context.Context, limit int) ([]ArticleHeader, error) // TODO add shift
	GetArticle(ctx context.Context, Id uuid.UUID) (*Article, error)
	CreateArticle(ctx context.Context, a *Article) error
	DeleteArticle(ctx context.Context, Id uuid.UUID) error
	io.Closer
}

package api

import (
	"io"

	"github.com/google/uuid"
)

type ArticleContent struct {
	Text string // Long string so I will store content as pointer
}

type Article struct {
	Id      uuid.UUID
	Title   string
	Context *ArticleContent
}

type ArticlesService interface {
	ListArticles() []Article
	io.Closer
}

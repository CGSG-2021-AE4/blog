package app

import "webapp/api"

type ArticlesService struct {
}

func NewArticlesService() *ArticlesService {
	return &ArticlesService{}
}

func (svc *ArticlesService) ListArticles() []api.Article {
	return []api.Article{
		{Title: "First", Context: nil},
		{Title: "Second", Context: nil},
		{Title: "Third", Context: nil},
	}
}

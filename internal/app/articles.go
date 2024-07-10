package app

import "github.com/CGSG-2021-AE4/blog/api"

type ArticlesService struct {
	Domain string
}

func NewArticlesService(domain string) *ArticlesService {
	return &ArticlesService{
		Domain: domain,
	}
}

func (svc *ArticlesService) ListArticles() []api.Article {
	return []api.Article{
		{Title: "First", Context: nil},
		{Title: "Second", Context: nil},
		{Title: "Third", Context: nil},
	}
}

func (svc *ArticlesService) Close() error {
	return nil
}

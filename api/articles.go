package api

type ArticleContent struct {
	Text string // Long string so I will store content as pointer
}

type Article struct {
	Title   string
	Context *ArticleContent
}

type ArticlesService interface {
	ListArticles() []Article
}

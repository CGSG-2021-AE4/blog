package articles

import (
	"net/http"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/api/router"
)

type ArticlesRouter struct {
	svc     api.ArticlesService
	userSvc api.UserService
}

func NewRouter(svc api.ArticlesService, userSvc api.UserService) router.Router {
	return &ArticlesRouter{
		svc:     svc,
		userSvc: userSvc,
	}
}

func (ar *ArticlesRouter) Routes() []router.Route {
	return []router.Route{
		// Pages
		{Method: http.MethodGet, Path: "/", Handler: router.ScriptPageHandler("index")},
		{Method: http.MethodGet, Path: "/article", Handler: router.ScriptPageHandler("article")},
		// Auth-required requests
		{Method: http.MethodGet, Path: "/article/edit", Handler: router.ScriptPageHandler("article_edit")},
		{Method: http.MethodGet, Path: "/article/create", Handler: router.ScriptPageHandler("article_create")}, // Now it is just setup name but later may be add tags or theme etc.

		// API
		{Method: http.MethodGet, Path: "/api/article/header", Handler: getArticleHandler(ar.svc)},          // Get article header
		{Method: http.MethodGet, Path: "/api/article/list", Handler: listArticlesHandler(ar.svc)},          // List some articles
		{Method: http.MethodGet, Path: "/api/article/contentHTML", Handler: getContentHTMLHandler(ar.svc)}, // Get article content converted to HTML
		{Method: http.MethodGet, Path: "/api/article/content", Handler: getContentHandler(ar.svc)},         // Get article content in MD format
		// Auth-required requests
		{Method: http.MethodPost, Path: "/api/article/create", Handler: createArticleHandler(ar.svc, ar.userSvc)},
		{Method: http.MethodPost, Path: "/api/article/edit", Handler: editArticleHandler(ar.svc)},
		{Method: http.MethodPost, Path: "/api/article/delete", Handler: deleteArticleHandler(ar.svc)},
	}
}

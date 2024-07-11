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
		{Method: http.MethodGet, Path: "/", Handler: router.ScriptPageHandler("index")},

		{Method: http.MethodGet, Path: "/api/article/list", Handler: listArticlesHandler(ar.svc)},
	}
}

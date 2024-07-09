package articles

import (
	"net/http"
	"webapp/api/router"
	"webapp/internal/app"
)

type ArticlesRouter struct {
	svc *app.ArticlesService
}

func NewRouter(svc *app.ArticlesService) router.Router {
	return &ArticlesRouter{
		svc: svc,
	}
}

func (ar *ArticlesRouter) Routes() []router.Route {
	return []router.Route{
		{Method: http.MethodGet, Path: "/", Handler: MainPageHandler(ar.svc)},
	}
}

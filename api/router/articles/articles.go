package articles

import (
	"net/http"
	"webapp/api"
	"webapp/api/router"
)

type ArticlesRouter struct {
	domain  string
	svc     api.ArticlesService
	userSvc api.UserService
}

func NewRouter(domain string, svc api.ArticlesService, userSvc api.UserService) router.Router {
	return &ArticlesRouter{
		domain:  domain,
		svc:     svc,
		userSvc: userSvc,
	}
}

func (ar *ArticlesRouter) Routes() []router.Route {
	return []router.Route{
		{Method: http.MethodGet, Path: "/", Handler: mainPageHandler(ar.domain, ar.svc, ar.userSvc)},
	}
}

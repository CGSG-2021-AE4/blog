package users

import (
	"net/http"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/api/router"
)

type UsersRouter struct {
	userSvc api.UserService
}

func NewRouter(userSvc api.UserService) router.Router {
	return &UsersRouter{
		userSvc: userSvc,
	}
}

func (ar *UsersRouter) Routes() []router.Route {
	return []router.Route{
		{Method: http.MethodGet, Path: "/login", Handler: router.ScriptPageHandler("login")},
		{Method: http.MethodGet, Path: "/signup", Handler: router.ScriptPageHandler("signup")},
		{Method: http.MethodGet, Path: "/account", Handler: router.ScriptPageHandler("account")},

		{Method: http.MethodPost, Path: "/api/user/login", Handler: loginHandler(ar.userSvc)},
		{Method: http.MethodPost, Path: "/api/user/reg", Handler: registerHandler(ar.userSvc)},
		{Method: http.MethodPost, Path: "/api/user/delete", Handler: deleteHandler(ar.userSvc)},
		{Method: http.MethodPost, Path: "/api/user/getPrivate", Handler: getUserPrivateHandler(ar.userSvc)},
		{Method: http.MethodPost, Path: "/api/user/getPublic", Handler: getUserPublicHandler(ar.userSvc)},
	}
}

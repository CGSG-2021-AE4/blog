package router

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

type Router interface {
	Routes() []Route
}

type RouterFunc func() []Route

func (rf RouterFunc) Routes() []Route {
	return rf()
}

type Routers struct {
	Rs []Router
}

// Collect all routers - therefore we can make even tree structure!!! cool
func (rs Routers) Routes() []Route {
	outRs := []Route{}

	for _, r := range rs.Rs {
		outRs = append(outRs, r.Routes()...)
	}
	return outRs
}

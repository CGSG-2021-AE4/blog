package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

type ErrorResp struct {
	Err string `json:"err"`
}

type TextResp struct {
	Text string `json:"text"`
}

func ScriptPageHandler(script string) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Origin": origin,
			"Title":  "AE4 blog",
			"Script": script,
		})
	}
}

package articles

import (
	"net/http"
	"strconv"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/api/router"
	"github.com/gin-gonic/gin"
)

func listArticlesHandler(as api.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		limit := 10
		if l, err := strconv.Atoi(c.Request.URL.Query().Get("limit")); err == nil {
			limit = l
		}

		articles, err := as.ListArticles(c, limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		c.JSON(http.StatusOK, articles)
	}
}

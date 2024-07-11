package articles

import (
	"net/http"
	"strconv"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/api/router"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func getArticleHandler(as api.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Request.URL.Query().Get("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "no id presented"})
			return
		}
		id, err := uuid.FromBytes([]byte(idStr))
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		a, err := as.GetArticle(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		c.JSON(http.StatusOK, a)
	}
}

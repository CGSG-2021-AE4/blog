package articles

import (
	"net/http"
	"webapp/internal/app"

	"github.com/gin-gonic/gin"
)

func MainPageHandler(svc *app.ArticlesService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "main.html", gin.H{
			"Title": "Title",
			"Body":  "Main page body",
		})
	}
}

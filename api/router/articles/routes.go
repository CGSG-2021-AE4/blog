package articles

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/api/router"
	"github.com/CGSG-2021-AE4/blog/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func listArticlesHandler(as api.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("List articles")
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
		log.Println("Get article")
		idStr := c.Request.URL.Query().Get("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "no id presented"})
			return
		}
		id, err := uuid.Parse(idStr)
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

func getContentHandler(as api.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Get article content")
		idStr := c.Request.URL.Query().Get("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "no id presented"})
			return
		}
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		content, err := as.GetContent(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		log.Println(string(content))
		c.JSON(http.StatusOK, router.TextResp{Text: string(content)})
	}
}

func getContentHTMLHandler(as api.ArticlesService) gin.HandlerFunc {

	return func(c *gin.Context) {
		mdParser := parser.NewWithExtensions(parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock)
		mdRender := html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags | html.HrefTargetBlank})
		log.Println("Get content HTML")
		idStr := c.Request.URL.Query().Get("id")
		if idStr == "" {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "no id presented"})
			return
		}
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		contentMD, err := as.GetContent(c, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		log.Println(string(contentMD))
		// TODO convert to html
		doc := mdParser.Parse(contentMD)
		log.Println(doc)
		contentHTML := markdown.Render(doc, mdRender)
		c.JSON(http.StatusOK, router.TextResp{Text: string(contentHTML)})
	}
}

type createArticleReq struct {
	Title string `json:"title"`
}
type createArticleResp struct {
	Id uuid.UUID `json:"id"`
}

func createArticleHandler(as api.ArticlesService, us api.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Create article")

		// Authorization check
		if c.Keys["authorized"] != "true" {
			log.Println("Not authorized error", c.Keys["authErr"])
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "not authorized"})
			return
		}

		// Decode request
		var req createArticleReq
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		log.Println(req)
		user, err := us.GetUser(c, c.Keys["authId"].(uuid.UUID))
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		descr := types.ArticleDescr{
			Title:          req.Title,
			AuthorId:       user.Id,
			AuthorUsername: user.Username,
		}

		// Process request
		id, err := as.CreateArticle(c, descr)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		c.JSON(http.StatusOK, createArticleResp{Id: id})
	}
}

type editArticleReq struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}
type editArticleResp struct {
	Msg string `json:"msg"`
}

func editArticleHandler(as api.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Edit article")

		// Authorization check
		if c.Keys["authorized"] != "true" {
			log.Println("Not authorized error", c.Keys["authErr"])
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "not authorized"})
			return
		}

		// Decode request
		var req editArticleReq
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("decode req: %w", err).Error()})
			return
		}
		// Check rights
		log.Println(req)
		a, err := as.GetArticle(c, req.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("get article: %w", err).Error()})
			return
		}
		if a.AuthorId != c.Keys["authId"].(uuid.UUID) {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "authorized and author users do not match"})
			return
		}

		// Process request
		if err := as.EditContent(c, a.ContentId, []byte(req.Content)); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("edit content: %w", err).Error()})
			return
		}
		a.ArticleDescr.Title = req.Title
		if err := as.EditArticle(c, a.Id, a.ArticleDescr); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: fmt.Errorf("edit article: %w", err).Error()})
			return
		}
		c.JSON(http.StatusOK, editArticleResp{Msg: "Edit complete"})
	}
}

type deleteArticleReq struct {
	Id uuid.UUID `json:"id"`
}
type deleteArticleResp struct {
	Msg string `json:"msg"`
}

func deleteArticleHandler(as api.ArticlesService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Delete article")

		// Authorization check
		if c.Keys["authorized"] != "true" {
			log.Println("Not authorized error", c.Keys["authErr"])
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "not authorized"})
			return
		}

		// Decode request
		var req deleteArticleReq
		if err := json.NewDecoder(c.Request.Body).Decode(&req); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		a, err := as.GetArticle(c, req.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		// Check rights
		if a.AuthorId != c.Keys["authId"].(uuid.UUID) {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: "authorized and author users do not match"})
			return
		}
		if err := as.DeleteArticle(c, a.Id); err != nil {
			c.JSON(http.StatusBadRequest, router.ErrorResp{Err: err.Error()})
			return
		}
		c.JSON(http.StatusOK, deleteArticleResp{Msg: "Delete complete"})
	}
}

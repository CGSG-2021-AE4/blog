package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/CGSG-2021-AE4/blog/api"
	"github.com/CGSG-2021-AE4/blog/api/router"
	"github.com/CGSG-2021-AE4/blog/api/router/articles"
	"github.com/CGSG-2021-AE4/blog/api/router/midleware"
	"github.com/CGSG-2021-AE4/blog/api/router/users"
	"github.com/CGSG-2021-AE4/blog/internal/app"
	"github.com/CGSG-2021-AE4/blog/internal/db/json"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("CGSG forever!!!")

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	// Parse flags

	conf := NewConfigFromFlags()

	if err := mainRun(ctx, conf); err != nil {
		log.Println("Finish main with error: ", err.Error())
	}
	log.Println("END")
}

func mainRun(ctx context.Context, conf Config) error {
	// Get router from services

	// Create stores
	userStore, err := json.NewUserStore(conf.UserStoreFilename)
	if err != nil {
		return err
	}
	articleStore, err := json.NewArticleStore(conf.ArticleStoreFilename)
	if err != nil {
		return err
	}

	// Create services
	articlesSvc := app.NewArticlesService(conf.Domain)
	userSvc := app.NewUserService(conf.UserSvcSecret, time.Duration(conf.UserTokenExpTimeout), userStore)

	routers := router.Routers{Rs: []router.Router{
		articles.NewRouter(conf.Domain, articlesSvc, userSvc),
		users.NewRouter(userSvc),
	}}

	midleware := []gin.HandlerFunc{
		midleware.AuthHandler(userSvc),
	}

	// Create server
	server := api.ApiServer{
		Addr:            conf.Domain,
		Midleware:       midleware,
		Router:          routers,
		CertFilename:    conf.CertFilename,
		PrivKeyFilename: conf.PrivKeyFilename,
	}

	// Run server
	if err := server.Start(ctx); err != nil {
		return nil
	}

	// Close services
	if err := userSvc.Close(); err != nil {
		return err
	}
	if err := articlesSvc.Close(); err != nil {
		return err
	}

	// Close stores
	if err := articleStore.Close(); err != nil {
		return err
	}
	if err := userStore.Close(); err != nil {
		return err
	}
	return nil
}

/* Registration testing * /


a := db.Article{
		Header: db.ArticleHeader{
			Id:    uuid.New(),
			Title: "Test title",
		},
		Content: &db.ArticleContent{
			Text: "Bla bla bla bla",
		},
	}
	if err := articleStore.CreateArticle(ctx, &a); err != nil {
		return err
	}



u := api.User{
	Id:       uuid.New(),
	Username: "yotia",
	Password: "cgsgforever",
	Email:    "what@the.fuck",
}
if err := userSvc.Register(ctx, &u); err != nil {
	return err
}
token, err := userSvc.Login(ctx, "yotia", "cgsgforever")
if err != nil {
	return err
}
claims, err := userSvc.ValidateToken(ctx, token)
if err != nil {
	return err
}
log.Println(claims)
/**/

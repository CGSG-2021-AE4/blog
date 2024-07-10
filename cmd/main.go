package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
	"webapp/api"
	"webapp/api/router"
	"webapp/api/router/articles"
	"webapp/internal/app"
	"webapp/internal/db/json"
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

	// Create services
	articlesSvc := app.NewArticlesService(conf.Domain)
	userSvc := app.NewUserService(conf.UserSvcSecret, time.Duration(conf.UserTokenExpTimeout), userStore)

	routers := router.Routers{Rs: []router.Router{
		articles.NewRouter(conf.Domain, articlesSvc, userSvc),
	}}

	// Create server
	server := api.ApiServer{
		Addr:            conf.Domain,
		Router:          routers,
		CertFilename:    conf.CertFilename,
		PrivKeyFilename: conf.PrivKeyFilename,
	}

	/* Registration testing * /
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

	// Run server
	if err := server.Start(ctx); err != nil {
		return nil
	}

	// Closing
	if err := userSvc.Close(); err != nil {
		return err
	}
	if err := articlesSvc.Close(); err != nil {
		return err
	}
	if err := userStore.Close(); err != nil {
		return err
	}
	return nil
}

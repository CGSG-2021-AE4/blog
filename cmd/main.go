package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"webapp/api"
	"webapp/api/router"
	"webapp/api/router/articles"
	"webapp/internal/app"
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

	articlesSvc := app.NewArticlesService()

	routers := router.Routers{Rs: []router.Router{
		articles.NewRouter(articlesSvc),
	}}

	// Create server
	server := api.ApiServer{
		Addr:            conf.Domain,
		Router:          routers,
		CertFilename:    conf.CertFilename,
		PrivKeyFilename: conf.PrivKeyFilename,
	}

	// Run server
	return server.Start(ctx)
}

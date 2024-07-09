package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"webapp/api/router"

	"github.com/gin-gonic/gin"
)

// TODO

type ApiServer struct {
	Addr   string
	Router router.Router

	stopCtx    context.CancelFunc
	httpServer http.Server
}

func (s *ApiServer) Start(ctx context.Context) error {
	ctx, s.stopCtx = context.WithCancel(ctx)

	rt := gin.New()

	// Load static files and templates
	rt.Delims("{", "}")
	rt.Static("bin/", "./web/bin")
	rt.LoadHTMLFiles("./web/templates/main.html")

	for _, r := range s.Router.Routes() {
		rt.Handle(r.Method, r.Path, r.Handler)
	}

	s.httpServer = http.Server{Addr: s.Addr, Handler: rt}

	// Handle context done
	go func() {
		<-ctx.Done()
		if err := s.Stop(); err != nil {
			log.Println("Failed to stop server:", err.Error())
		}
	}()

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *ApiServer) Stop() error {
	s.stopCtx()
	// Stop http server
	if err := s.httpServer.Close(); err != nil {
		return err
	}
	return nil
}

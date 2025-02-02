package api

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net/http"

	"github.com/CGSG-2021-AE4/blog/api/router"
	"github.com/CGSG-2021-AE4/blog/pkg/cg"
	"github.com/CGSG-2021-AE4/blog/pkg/cg/sscg"

	"github.com/gin-gonic/gin"
)

// TODO

type ApiServer struct {
	Addr            string // Domain addr
	Midleware       []gin.HandlerFunc
	Router          router.Router
	CertFilename    string
	PrivKeyFilename string

	stopCtx    context.CancelFunc
	httpServer http.Server
}

func (s *ApiServer) Start(ctx context.Context) error {
	ctx, s.stopCtx = context.WithCancel(ctx)

	// Check certificate and private key
	if err := cg.CheckCert(s.CertFilename); err != nil {
		log.Println("Certificate is invalid")
		log.Println("Regenerate certificate...")
		if err := sscg.Gen(s.CertFilename, s.PrivKeyFilename); err != nil {
			return err
		}
	}

	rt := gin.New()

	// Load static files and templates
	rt.Delims("{", "}")
	rt.Static("bin/", "./web/bin")
	rt.LoadHTMLFiles("./web/templates/index.html")

	for _, f := range s.Midleware {
		rt.Use(f)
	}

	for _, r := range s.Router.Routes() {
		rt.Handle(r.Method, r.Path, r.Handler)
	}

	s.httpServer = http.Server{
		Addr:    s.Addr,
		Handler: rt,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
	}

	// Handle context done
	go func() {
		<-ctx.Done()
		if err := s.Stop(); err != nil {
			log.Println("Failed to stop server:", err.Error())
		}
	}()

	// if err := s.httpServer.ListenAndServeTLS(s.CertFilename, s.PrivKeyFilename); err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 	return err
	// }
	// Without https
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

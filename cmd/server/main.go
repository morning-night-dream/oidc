package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/morning-night-dream/oidc/handler/idp"
	"github.com/morning-night-dream/oidc/handler/op"
	"github.com/morning-night-dream/oidc/handler/rp"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func main() {
	srv := NewServer("1234", NewHandler())

	srv.Run()
}

type Server struct {
	*http.Server
}

func NewServer(
	port string,
	handler http.Handler,
) *Server {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler,
	}

	return &Server{
		Server: srv,
	}
}

func (srv *Server) Run() {
	log.Printf("server started on %s", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("server closed with error: %v", err)

			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("failed to gracefully shutdown: %v", err)
	}

	log.Printf("server shutdown")
}

var _ openapi.ServerInterface = (*Handler)(nil)

type Handler struct{}

func NewHandler() http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {})

	hdl := openapi.HandlerWithOptions(&Handler{}, openapi.ChiServerOptions{
		BaseRouter: router,
	})

	return hdl
}

func (hdl *Handler) IdpSignUp(
	w http.ResponseWriter,
	r *http.Request,
) {
	idp.SignUp(w, r)
}

func (hdl *Handler) IdpSignIn(
	w http.ResponseWriter,
	r *http.Request,
) {
	idp.SignIn(w, r)
}

func (hdl *Handler) OpOpenIDConfiguration(
	w http.ResponseWriter,
	r *http.Request,
) {
	op.OpenIDConfiguration(w, r)
}

func (hdl *Handler) OpToken(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpTokenParams,
) {
}

func (hdl *Handler) OpAuthorize(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpAuthorizeParams,
) {
	op.Auth(w, r)
}

func (hdl *Handler) RpLogin(
	w http.ResponseWriter,
	r *http.Request,
) {
	rp.Login(w, r)
}

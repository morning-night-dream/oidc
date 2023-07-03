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
	"github.com/morning-night-dream/oidc/cache"
	"github.com/morning-night-dream/oidc/handler/idp"
	"github.com/morning-night-dream/oidc/handler/op"
	"github.com/morning-night-dream/oidc/handler/rp"
	"github.com/morning-night-dream/oidc/middleware"
	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func main() {
	user := cache.New[model.User]()

	password, err := model.PasswordEncrypt("password")
	if err != nil {
		panic(err)
	}

	// NOTE: テストユーザーを登録
	user.Set("username", model.User{
		ID:       "id",
		Username: "username",
		Password: password,
	})

	accessToken := cache.New[model.AccessToken]()

	refreshToken := cache.New[model.RefreshToken]()

	idToken := cache.New[model.IDToken]()

	idp := &idp.IdP{
		UserCache: user,
	}

	rp := &rp.RP{
		ClientID:    "morning-night-dream",
		RedirectURI: "http://localhost:1234/rp/callback",
		Scopes:      []string{"openid"},
		AuthURL:     "http://localhost:1234/op/authorize",
		TokenURL:    "http://localhost:1234/op/token",
		UserInfoURL: "http://localhost:1234/op/userinfo",
	}

	op := &op.OP{
		AllowClientID:        "morning-night-dream",
		AllowRedirectURI:     "http://localhost:1234/rp/callback",
		AuthorizeParamsCache: cache.New[openapi.OpAuthorizeParams](),
		UserCache:            user,
		LoggedInUserCache:    cache.New[model.User](),
		AccessTokenCache:     accessToken,
		RefreshTokenCache:    refreshToken,
		IDTokenCache:         idToken,
		Issuer:               "http://localhost:1234",
	}

	srv := NewServer("1234", NewHandler(idp, rp, op))

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

type Handler struct {
	IdP *idp.IdP
	RP  *rp.RP
	OP  *op.OP
}

func NewHandler(
	idp *idp.IdP,
	rp *rp.RP,
	op *op.OP,
) http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {})

	// NOTE: 動作確認しやすくするための実装
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		url := "http://localhost:1234/rp/login"

		http.Redirect(w, r, url, http.StatusFound)
	})

	hdl := openapi.HandlerWithOptions(
		&Handler{
			IdP: idp,
			RP:  rp,
			OP:  op,
		},
		openapi.ChiServerOptions{
			BaseRouter: router,
			Middlewares: []openapi.MiddlewareFunc{
				middleware.NewAccessLog().Handle,
				middleware.NewCORS(),
			},
		})

	return hdl
}

func (hdl *Handler) IdpSignup(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.IdP.Signup(w, r)
}

func (hdl *Handler) IdpSignin(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.IdP.Signin(w, r)
}

func (hdl *Handler) OpOpenIDConfiguration(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.OP.OpenIDConfiguration(w, r)
}

func (hdl *Handler) OpAuthorize(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpAuthorizeParams,
) {
	hdl.OP.Authorize(w, r, params)
}

func (hdl *Handler) OpLoginView(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpLoginViewParams,
) {
	hdl.OP.LoginView(w, r, params)
}

func (hdl *Handler) OpLogin(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.OP.Login(w, r)
}

func (hdl *Handler) OpCallback(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpCallbackParams,
) {
	hdl.OP.Callback(w, r, params)
}

func (hdl *Handler) OpUserinfo(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.OP.Userinfo(w, r)
}

func (hdl *Handler) OpToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.OP.Token(w, r)
}

func (hdl *Handler) RpLogin(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.RP.Login(w, r)
}

func (hdl *Handler) RpCallback(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.RpCallbackParams,
) {
	hdl.RP.Callback(w, r, params)
}

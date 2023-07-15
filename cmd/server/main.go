package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/morning-night-dream/oidc/cache"
	"github.com/morning-night-dream/oidc/handler/idp"
	"github.com/morning-night-dream/oidc/handler/op"
	"github.com/morning-night-dream/oidc/handler/rp"
	"github.com/morning-night-dream/oidc/middleware"
	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	selfURL := os.Getenv("SELF_URL")
	if selfURL == "" {
		selfURL = fmt.Sprintf("http://localhost:%s", port)
	}

	user := cache.New[model.User]()

	password, err := model.PasswordEncrypt("password")
	if err != nil {
		panic(err)
	}

	// NOTE: テストユーザーを登録
	user.Set("username", model.User{
		ID:       uuid.NewString(),
		Username: "username",
		Password: password,
	})

	accessToken := cache.New[model.User]()

	refreshToken := cache.New[model.User]()

	idToken := cache.New[model.User]()

	idp := &idp.IdP{
		UserCache: user,
	}

	rp := &rp.RP{
		ClientID:    "morning-night-dream",
		RedirectURI: fmt.Sprintf("%s/rp/callback", selfURL),
		Scopes:      []string{"openid"},
		AuthURL:     fmt.Sprintf("%s/op/authorize", selfURL),
		TokenURL:    fmt.Sprintf("%s/op/token", selfURL),
		UserInfoURL: fmt.Sprintf("%s/op/userinfo", selfURL),
	}

	reader := rand.Reader

	bitSize := 2048

	key, err := rsa.GenerateKey(reader, bitSize)
	if err != nil {
		panic(err)
	}

	op := &op.OP{
		AllowClientID:        "morning-night-dream",
		AllowRedirectURI:     fmt.Sprintf("%s/rp/callback", selfURL),
		AuthorizeParamsCache: cache.New[openapi.OpAuthorizeParams](),
		PrivateKey:           key,
		UserCache:            user,
		LoggedInUserCache:    cache.New[model.User](),
		AccessTokenCache:     accessToken,
		RefreshTokenCache:    refreshToken,
		IDTokenCache:         idToken,
		Issuer:               selfURL,
	}

	srv := NewServer(port, NewHandler(selfURL, idp, rp, op))

	log.Printf("server started on %s", selfURL)

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
	selfURL string,
	idp *idp.IdP,
	rp *rp.RP,
	op *op.OP,
) http.Handler {
	router := chi.NewRouter()

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {})

	// NOTE: 動作確認しやすくするための実装
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s/rp/login", selfURL)

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

func (hdl *Handler) OpCerts(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.OP.Certs(w, r)
}

func (hdl *Handler) OpRevoke(
	w http.ResponseWriter,
	r *http.Request,
) {
	hdl.OP.Revoke(w, r)
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

// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Sign In
	// (POST /id/signin)
	SignIn(w http.ResponseWriter, r *http.Request)
	// Sign Up
	// (POST /id/signup)
	SignUp(w http.ResponseWriter, r *http.Request)
	// OpenID Provider Configuration
	// (GET /op/.well-known/openid-configuration)
	OpenIDConfiguration(w http.ResponseWriter, r *http.Request)
	// Authentication Request
	// (GET /op/authorize)
	Authorize(w http.ResponseWriter, r *http.Request, params AuthorizeParams)
	// Token Request
	// (POST /op/token)
	Token(w http.ResponseWriter, r *http.Request, params TokenParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// SignIn operation middleware
func (siw *ServerInterfaceWrapper) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SignIn(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// SignUp operation middleware
func (siw *ServerInterfaceWrapper) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SignUp(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpenIDConfiguration operation middleware
func (siw *ServerInterfaceWrapper) OpenIDConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OpenIDConfiguration(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Authorize operation middleware
func (siw *ServerInterfaceWrapper) Authorize(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AuthorizeParams

	// ------------- Optional query parameter "response_type" -------------

	err = runtime.BindQueryParameter("form", true, false, "response_type", r.URL.Query(), &params.ResponseType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "response_type", Err: err})
		return
	}

	// ------------- Optional query parameter "scope" -------------

	err = runtime.BindQueryParameter("form", true, false, "scope", r.URL.Query(), &params.Scope)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "scope", Err: err})
		return
	}

	// ------------- Optional query parameter "client_id" -------------

	err = runtime.BindQueryParameter("form", true, false, "client_id", r.URL.Query(), &params.ClientId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "client_id", Err: err})
		return
	}

	// ------------- Optional query parameter "state" -------------

	err = runtime.BindQueryParameter("form", true, false, "state", r.URL.Query(), &params.State)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "state", Err: err})
		return
	}

	// ------------- Optional query parameter "redirect_uri" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirect_uri", r.URL.Query(), &params.RedirectUri)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "redirect_uri", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Authorize(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// Token operation middleware
func (siw *ServerInterfaceWrapper) Token(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params TokenParams

	// ------------- Optional query parameter "grant_type" -------------

	err = runtime.BindQueryParameter("form", true, false, "grant_type", r.URL.Query(), &params.GrantType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "grant_type", Err: err})
		return
	}

	// ------------- Optional query parameter "code" -------------

	err = runtime.BindQueryParameter("form", true, false, "code", r.URL.Query(), &params.Code)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "code", Err: err})
		return
	}

	// ------------- Optional query parameter "redirect_uri" -------------

	err = runtime.BindQueryParameter("form", true, false, "redirect_uri", r.URL.Query(), &params.RedirectUri)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "redirect_uri", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Token(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/id/signin", wrapper.SignIn)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/id/signup", wrapper.SignUp)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/op/.well-known/openid-configuration", wrapper.OpenIDConfiguration)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/op/authorize", wrapper.Authorize)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/op/token", wrapper.Token)
	})

	return r
}

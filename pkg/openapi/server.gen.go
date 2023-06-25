// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Sign In
	// (POST /id/signin)
	IdpSignIn(w http.ResponseWriter, r *http.Request)
	// Sign Up
	// (POST /idp/signup)
	IdpSignUp(w http.ResponseWriter, r *http.Request)
	// OpenID Provider Configuration
	// (GET /op/.well-known/openid-configuration)
	OpOpenIDConfiguration(w http.ResponseWriter, r *http.Request)
	// Authentication Request
	// (GET /op/authorize)
	OpAuthorize(w http.ResponseWriter, r *http.Request, params OpAuthorizeParams)
	// OP Callback
	// (GET /op/callback)
	OpCallback(w http.ResponseWriter, r *http.Request, params OpCallbackParams)
	// OP Login
	// (POST /op/login)
	OpLogin(w http.ResponseWriter, r *http.Request)
	// OP Login
	// (GET /op/login/view)
	OpLoginView(w http.ResponseWriter, r *http.Request, params OpLoginViewParams)
	// Token Request
	// (POST /op/token)
	OpToken(w http.ResponseWriter, r *http.Request, params OpTokenParams)
	// UserInfo Request
	// (GET /op/userinfo)
	OpUserinfo(w http.ResponseWriter, r *http.Request)
	// RP Callback
	// (GET /rp/callback)
	RpCallback(w http.ResponseWriter, r *http.Request, params RpCallbackParams)
	// RP Login
	// (GET /rp/login)
	RpLogin(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// IdpSignIn operation middleware
func (siw *ServerInterfaceWrapper) IdpSignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.IdpSignIn(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// IdpSignUp operation middleware
func (siw *ServerInterfaceWrapper) IdpSignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.IdpSignUp(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpOpenIDConfiguration operation middleware
func (siw *ServerInterfaceWrapper) OpOpenIDConfiguration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OpOpenIDConfiguration(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpAuthorize operation middleware
func (siw *ServerInterfaceWrapper) OpAuthorize(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OpAuthorizeParams

	// ------------- Required query parameter "response_type" -------------

	if paramValue := r.URL.Query().Get("response_type"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "response_type"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "response_type", r.URL.Query(), &params.ResponseType)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "response_type", Err: err})
		return
	}

	// ------------- Required query parameter "scope" -------------

	if paramValue := r.URL.Query().Get("scope"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "scope"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "scope", r.URL.Query(), &params.Scope)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "scope", Err: err})
		return
	}

	// ------------- Required query parameter "client_id" -------------

	if paramValue := r.URL.Query().Get("client_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "client_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "client_id", r.URL.Query(), &params.ClientId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "client_id", Err: err})
		return
	}

	// ------------- Required query parameter "redirect_uri" -------------

	if paramValue := r.URL.Query().Get("redirect_uri"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "redirect_uri"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "redirect_uri", r.URL.Query(), &params.RedirectUri)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "redirect_uri", Err: err})
		return
	}

	// ------------- Optional query parameter "state" -------------

	err = runtime.BindQueryParameter("form", true, false, "state", r.URL.Query(), &params.State)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "state", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OpAuthorize(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpCallback operation middleware
func (siw *ServerInterfaceWrapper) OpCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OpCallbackParams

	// ------------- Required query parameter "id" -------------

	if paramValue := r.URL.Query().Get("id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "id", r.URL.Query(), &params.Id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OpCallback(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpLogin operation middleware
func (siw *ServerInterfaceWrapper) OpLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OpLogin(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpLoginView operation middleware
func (siw *ServerInterfaceWrapper) OpLoginView(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OpLoginViewParams

	// ------------- Required query parameter "auth_request_id" -------------

	if paramValue := r.URL.Query().Get("auth_request_id"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "auth_request_id"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "auth_request_id", r.URL.Query(), &params.AuthRequestId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "auth_request_id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OpLoginView(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpToken operation middleware
func (siw *ServerInterfaceWrapper) OpToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params OpTokenParams

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
		siw.Handler.OpToken(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// OpUserinfo operation middleware
func (siw *ServerInterfaceWrapper) OpUserinfo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerScopes, []string{})

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.OpUserinfo(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RpCallback operation middleware
func (siw *ServerInterfaceWrapper) RpCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params RpCallbackParams

	// ------------- Required query parameter "code" -------------

	if paramValue := r.URL.Query().Get("code"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "code"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "code", r.URL.Query(), &params.Code)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "code", Err: err})
		return
	}

	// ------------- Required query parameter "state" -------------

	if paramValue := r.URL.Query().Get("state"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "state"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "state", r.URL.Query(), &params.State)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "state", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RpCallback(w, r, params)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// RpLogin operation middleware
func (siw *ServerInterfaceWrapper) RpLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RpLogin(w, r)
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
		r.Post(options.BaseURL+"/id/signin", wrapper.IdpSignIn)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/idp/signup", wrapper.IdpSignUp)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/op/.well-known/openid-configuration", wrapper.OpOpenIDConfiguration)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/op/authorize", wrapper.OpAuthorize)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/op/callback", wrapper.OpCallback)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/op/login", wrapper.OpLogin)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/op/login/view", wrapper.OpLoginView)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/op/token", wrapper.OpToken)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/op/userinfo", wrapper.OpUserinfo)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/rp/callback", wrapper.RpCallback)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/rp/login", wrapper.RpLogin)
	})

	return r
}

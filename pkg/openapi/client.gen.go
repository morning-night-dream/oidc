// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// IdpSignIn request with any body
	IdpSignInWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	IdpSignIn(ctx context.Context, body IdpSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// IdpSignUp request with any body
	IdpSignUpWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	IdpSignUp(ctx context.Context, body IdpSignUpJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpOpenIDConfiguration request
	OpOpenIDConfiguration(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpAuthorize request
	OpAuthorize(ctx context.Context, params *OpAuthorizeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpToken request
	OpToken(ctx context.Context, params *OpTokenParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RpCallback request
	RpCallback(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RpLogin request
	RpLogin(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) IdpSignInWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSignInRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) IdpSignIn(ctx context.Context, body IdpSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSignInRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) IdpSignUpWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSignUpRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) IdpSignUp(ctx context.Context, body IdpSignUpJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSignUpRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpOpenIDConfiguration(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpOpenIDConfigurationRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpAuthorize(ctx context.Context, params *OpAuthorizeParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpAuthorizeRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpToken(ctx context.Context, params *OpTokenParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpTokenRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RpCallback(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRpCallbackRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RpLogin(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRpLoginRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewIdpSignInRequest calls the generic IdpSignIn builder with application/json body
func NewIdpSignInRequest(server string, body IdpSignInJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewIdpSignInRequestWithBody(server, "application/json", bodyReader)
}

// NewIdpSignInRequestWithBody generates requests for IdpSignIn with any type of body
func NewIdpSignInRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/id/signin")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewIdpSignUpRequest calls the generic IdpSignUp builder with application/json body
func NewIdpSignUpRequest(server string, body IdpSignUpJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewIdpSignUpRequestWithBody(server, "application/json", bodyReader)
}

// NewIdpSignUpRequestWithBody generates requests for IdpSignUp with any type of body
func NewIdpSignUpRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/idp/signup")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewOpOpenIDConfigurationRequest generates requests for OpOpenIDConfiguration
func NewOpOpenIDConfigurationRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/op/.well-known/openid-configuration")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewOpAuthorizeRequest generates requests for OpAuthorize
func NewOpAuthorizeRequest(server string, params *OpAuthorizeParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/op/authorize")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.ResponseType != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "response_type", runtime.ParamLocationQuery, *params.ResponseType); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Scope != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "scope", runtime.ParamLocationQuery, *params.Scope); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.ClientId != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "client_id", runtime.ParamLocationQuery, *params.ClientId); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.State != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "state", runtime.ParamLocationQuery, *params.State); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.RedirectUri != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "redirect_uri", runtime.ParamLocationQuery, *params.RedirectUri); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewOpTokenRequest generates requests for OpToken
func NewOpTokenRequest(server string, params *OpTokenParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/op/token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.GrantType != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "grant_type", runtime.ParamLocationQuery, *params.GrantType); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.Code != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "code", runtime.ParamLocationQuery, *params.Code); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		if params.RedirectUri != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "redirect_uri", runtime.ParamLocationQuery, *params.RedirectUri); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRpCallbackRequest generates requests for RpCallback
func NewRpCallbackRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/rp/callback")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRpLoginRequest generates requests for RpLogin
func NewRpLoginRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/rp/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// IdpSignIn request with any body
	IdpSignInWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSignInResponse, error)

	IdpSignInWithResponse(ctx context.Context, body IdpSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSignInResponse, error)

	// IdpSignUp request with any body
	IdpSignUpWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSignUpResponse, error)

	IdpSignUpWithResponse(ctx context.Context, body IdpSignUpJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSignUpResponse, error)

	// OpOpenIDConfiguration request
	OpOpenIDConfigurationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpOpenIDConfigurationResponse, error)

	// OpAuthorize request
	OpAuthorizeWithResponse(ctx context.Context, params *OpAuthorizeParams, reqEditors ...RequestEditorFn) (*OpAuthorizeResponse, error)

	// OpToken request
	OpTokenWithResponse(ctx context.Context, params *OpTokenParams, reqEditors ...RequestEditorFn) (*OpTokenResponse, error)

	// RpCallback request
	RpCallbackWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RpCallbackResponse, error)

	// RpLogin request
	RpLoginWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RpLoginResponse, error)
}

type IdpSignInResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r IdpSignInResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r IdpSignInResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type IdpSignUpResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r IdpSignUpResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r IdpSignUpResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type OpOpenIDConfigurationResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OPOpenIDConfigurationResponseSchema
}

// Status returns HTTPResponse.Status
func (r OpOpenIDConfigurationResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpOpenIDConfigurationResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type OpAuthorizeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r OpAuthorizeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpAuthorizeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type OpTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OPTokenResponseSchema
	JSON400      *struct {
		// Error error
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r OpTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RpCallbackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r RpCallbackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RpCallbackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RpLoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r RpLoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RpLoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// IdpSignInWithBodyWithResponse request with arbitrary body returning *IdpSignInResponse
func (c *ClientWithResponses) IdpSignInWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSignInResponse, error) {
	rsp, err := c.IdpSignInWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSignInResponse(rsp)
}

func (c *ClientWithResponses) IdpSignInWithResponse(ctx context.Context, body IdpSignInJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSignInResponse, error) {
	rsp, err := c.IdpSignIn(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSignInResponse(rsp)
}

// IdpSignUpWithBodyWithResponse request with arbitrary body returning *IdpSignUpResponse
func (c *ClientWithResponses) IdpSignUpWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSignUpResponse, error) {
	rsp, err := c.IdpSignUpWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSignUpResponse(rsp)
}

func (c *ClientWithResponses) IdpSignUpWithResponse(ctx context.Context, body IdpSignUpJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSignUpResponse, error) {
	rsp, err := c.IdpSignUp(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSignUpResponse(rsp)
}

// OpOpenIDConfigurationWithResponse request returning *OpOpenIDConfigurationResponse
func (c *ClientWithResponses) OpOpenIDConfigurationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpOpenIDConfigurationResponse, error) {
	rsp, err := c.OpOpenIDConfiguration(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpOpenIDConfigurationResponse(rsp)
}

// OpAuthorizeWithResponse request returning *OpAuthorizeResponse
func (c *ClientWithResponses) OpAuthorizeWithResponse(ctx context.Context, params *OpAuthorizeParams, reqEditors ...RequestEditorFn) (*OpAuthorizeResponse, error) {
	rsp, err := c.OpAuthorize(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpAuthorizeResponse(rsp)
}

// OpTokenWithResponse request returning *OpTokenResponse
func (c *ClientWithResponses) OpTokenWithResponse(ctx context.Context, params *OpTokenParams, reqEditors ...RequestEditorFn) (*OpTokenResponse, error) {
	rsp, err := c.OpToken(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpTokenResponse(rsp)
}

// RpCallbackWithResponse request returning *RpCallbackResponse
func (c *ClientWithResponses) RpCallbackWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RpCallbackResponse, error) {
	rsp, err := c.RpCallback(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRpCallbackResponse(rsp)
}

// RpLoginWithResponse request returning *RpLoginResponse
func (c *ClientWithResponses) RpLoginWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RpLoginResponse, error) {
	rsp, err := c.RpLogin(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRpLoginResponse(rsp)
}

// ParseIdpSignInResponse parses an HTTP response from a IdpSignInWithResponse call
func ParseIdpSignInResponse(rsp *http.Response) (*IdpSignInResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &IdpSignInResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseIdpSignUpResponse parses an HTTP response from a IdpSignUpWithResponse call
func ParseIdpSignUpResponse(rsp *http.Response) (*IdpSignUpResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &IdpSignUpResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseOpOpenIDConfigurationResponse parses an HTTP response from a OpOpenIDConfigurationWithResponse call
func ParseOpOpenIDConfigurationResponse(rsp *http.Response) (*OpOpenIDConfigurationResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpOpenIDConfigurationResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OPOpenIDConfigurationResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseOpAuthorizeResponse parses an HTTP response from a OpAuthorizeWithResponse call
func ParseOpAuthorizeResponse(rsp *http.Response) (*OpAuthorizeResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpAuthorizeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseOpTokenResponse parses an HTTP response from a OpTokenWithResponse call
func ParseOpTokenResponse(rsp *http.Response) (*OpTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OPTokenResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			// Error error
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseRpCallbackResponse parses an HTTP response from a RpCallbackWithResponse call
func ParseRpCallbackResponse(rsp *http.Response) (*RpCallbackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RpCallbackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseRpLoginResponse parses an HTTP response from a RpLoginWithResponse call
func ParseRpLoginResponse(rsp *http.Response) (*RpLoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RpLoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

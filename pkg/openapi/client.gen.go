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
	// IdpSignin request with any body
	IdpSigninWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	IdpSignin(ctx context.Context, body IdpSigninJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// IdpSignup request with any body
	IdpSignupWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	IdpSignup(ctx context.Context, body IdpSignupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpOpenIDConfiguration request
	OpOpenIDConfiguration(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpAuthorize request
	OpAuthorize(ctx context.Context, params *OpAuthorizeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpCallback request
	OpCallback(ctx context.Context, params *OpCallbackParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpLogin request
	OpLogin(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpLoginView request
	OpLoginView(ctx context.Context, params *OpLoginViewParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpToken request with any body
	OpTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	OpTokenWithFormdataBody(ctx context.Context, body OpTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// OpUserinfo request
	OpUserinfo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RpCallback request
	RpCallback(ctx context.Context, params *RpCallbackParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RpLogin request
	RpLogin(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) IdpSigninWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSigninRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) IdpSignin(ctx context.Context, body IdpSigninJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSigninRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) IdpSignupWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSignupRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) IdpSignup(ctx context.Context, body IdpSignupJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewIdpSignupRequest(c.Server, body)
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

func (c *Client) OpCallback(ctx context.Context, params *OpCallbackParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpCallbackRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpLogin(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpLoginRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpLoginView(ctx context.Context, params *OpLoginViewParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpLoginViewRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpTokenWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpTokenRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpTokenWithFormdataBody(ctx context.Context, body OpTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpTokenRequestWithFormdataBody(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) OpUserinfo(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewOpUserinfoRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RpCallback(ctx context.Context, params *RpCallbackParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRpCallbackRequest(c.Server, params)
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

// NewIdpSigninRequest calls the generic IdpSignin builder with application/json body
func NewIdpSigninRequest(server string, body IdpSigninJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewIdpSigninRequestWithBody(server, "application/json", bodyReader)
}

// NewIdpSigninRequestWithBody generates requests for IdpSignin with any type of body
func NewIdpSigninRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/idp/signin")
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

// NewIdpSignupRequest calls the generic IdpSignup builder with application/json body
func NewIdpSignupRequest(server string, body IdpSignupJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewIdpSignupRequestWithBody(server, "application/json", bodyReader)
}

// NewIdpSignupRequestWithBody generates requests for IdpSignup with any type of body
func NewIdpSignupRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
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

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "response_type", runtime.ParamLocationQuery, params.ResponseType); err != nil {
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

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "scope", runtime.ParamLocationQuery, params.Scope); err != nil {
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

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "client_id", runtime.ParamLocationQuery, params.ClientId); err != nil {
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

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "redirect_uri", runtime.ParamLocationQuery, params.RedirectUri); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewOpCallbackRequest generates requests for OpCallback
func NewOpCallbackRequest(server string, params *OpCallbackParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/op/callback")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "id", runtime.ParamLocationQuery, params.Id); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewOpLoginRequest generates requests for OpLogin
func NewOpLoginRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/op/login")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewOpLoginViewRequest generates requests for OpLoginView
func NewOpLoginViewRequest(server string, params *OpLoginViewParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/op/login/view")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "auth_request_id", runtime.ParamLocationQuery, params.AuthRequestId); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewOpTokenRequestWithFormdataBody calls the generic OpToken builder with application/x-www-form-urlencoded body
func NewOpTokenRequestWithFormdataBody(server string, body OpTokenFormdataRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	bodyStr, err := runtime.MarshalForm(body, nil)
	if err != nil {
		return nil, err
	}
	bodyReader = strings.NewReader(bodyStr.Encode())
	return NewOpTokenRequestWithBody(server, "application/x-www-form-urlencoded", bodyReader)
}

// NewOpTokenRequestWithBody generates requests for OpToken with any type of body
func NewOpTokenRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
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

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewOpUserinfoRequest generates requests for OpUserinfo
func NewOpUserinfoRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/op/userinfo")
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

// NewRpCallbackRequest generates requests for RpCallback
func NewRpCallbackRequest(server string, params *RpCallbackParams) (*http.Request, error) {
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

	if params != nil {
		queryValues := queryURL.Query()

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "code", runtime.ParamLocationQuery, params.Code); err != nil {
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

		if queryFrag, err := runtime.StyleParamWithLocation("form", true, "state", runtime.ParamLocationQuery, params.State); err != nil {
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

		queryURL.RawQuery = queryValues.Encode()
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
	// IdpSignin request with any body
	IdpSigninWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSigninResponse, error)

	IdpSigninWithResponse(ctx context.Context, body IdpSigninJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSigninResponse, error)

	// IdpSignup request with any body
	IdpSignupWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSignupResponse, error)

	IdpSignupWithResponse(ctx context.Context, body IdpSignupJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSignupResponse, error)

	// OpOpenIDConfiguration request
	OpOpenIDConfigurationWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpOpenIDConfigurationResponse, error)

	// OpAuthorize request
	OpAuthorizeWithResponse(ctx context.Context, params *OpAuthorizeParams, reqEditors ...RequestEditorFn) (*OpAuthorizeResponse, error)

	// OpCallback request
	OpCallbackWithResponse(ctx context.Context, params *OpCallbackParams, reqEditors ...RequestEditorFn) (*OpCallbackResponse, error)

	// OpLogin request
	OpLoginWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpLoginResponse, error)

	// OpLoginView request
	OpLoginViewWithResponse(ctx context.Context, params *OpLoginViewParams, reqEditors ...RequestEditorFn) (*OpLoginViewResponse, error)

	// OpToken request with any body
	OpTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*OpTokenResponse, error)

	OpTokenWithFormdataBodyWithResponse(ctx context.Context, body OpTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*OpTokenResponse, error)

	// OpUserinfo request
	OpUserinfoWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpUserinfoResponse, error)

	// RpCallback request
	RpCallbackWithResponse(ctx context.Context, params *RpCallbackParams, reqEditors ...RequestEditorFn) (*RpCallbackResponse, error)

	// RpLogin request
	RpLoginWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RpLoginResponse, error)
}

type IdpSigninResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r IdpSigninResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r IdpSigninResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type IdpSignupResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r IdpSignupResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r IdpSignupResponse) StatusCode() int {
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

type OpCallbackResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r OpCallbackResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpCallbackResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type OpLoginResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r OpLoginResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpLoginResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type OpLoginViewResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r OpLoginViewResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpLoginViewResponse) StatusCode() int {
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

type OpUserinfoResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *OPUserInfoResponseSchema
}

// Status returns HTTPResponse.Status
func (r OpUserinfoResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r OpUserinfoResponse) StatusCode() int {
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

// IdpSigninWithBodyWithResponse request with arbitrary body returning *IdpSigninResponse
func (c *ClientWithResponses) IdpSigninWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSigninResponse, error) {
	rsp, err := c.IdpSigninWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSigninResponse(rsp)
}

func (c *ClientWithResponses) IdpSigninWithResponse(ctx context.Context, body IdpSigninJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSigninResponse, error) {
	rsp, err := c.IdpSignin(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSigninResponse(rsp)
}

// IdpSignupWithBodyWithResponse request with arbitrary body returning *IdpSignupResponse
func (c *ClientWithResponses) IdpSignupWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*IdpSignupResponse, error) {
	rsp, err := c.IdpSignupWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSignupResponse(rsp)
}

func (c *ClientWithResponses) IdpSignupWithResponse(ctx context.Context, body IdpSignupJSONRequestBody, reqEditors ...RequestEditorFn) (*IdpSignupResponse, error) {
	rsp, err := c.IdpSignup(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseIdpSignupResponse(rsp)
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

// OpCallbackWithResponse request returning *OpCallbackResponse
func (c *ClientWithResponses) OpCallbackWithResponse(ctx context.Context, params *OpCallbackParams, reqEditors ...RequestEditorFn) (*OpCallbackResponse, error) {
	rsp, err := c.OpCallback(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpCallbackResponse(rsp)
}

// OpLoginWithResponse request returning *OpLoginResponse
func (c *ClientWithResponses) OpLoginWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpLoginResponse, error) {
	rsp, err := c.OpLogin(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpLoginResponse(rsp)
}

// OpLoginViewWithResponse request returning *OpLoginViewResponse
func (c *ClientWithResponses) OpLoginViewWithResponse(ctx context.Context, params *OpLoginViewParams, reqEditors ...RequestEditorFn) (*OpLoginViewResponse, error) {
	rsp, err := c.OpLoginView(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpLoginViewResponse(rsp)
}

// OpTokenWithBodyWithResponse request with arbitrary body returning *OpTokenResponse
func (c *ClientWithResponses) OpTokenWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*OpTokenResponse, error) {
	rsp, err := c.OpTokenWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpTokenResponse(rsp)
}

func (c *ClientWithResponses) OpTokenWithFormdataBodyWithResponse(ctx context.Context, body OpTokenFormdataRequestBody, reqEditors ...RequestEditorFn) (*OpTokenResponse, error) {
	rsp, err := c.OpTokenWithFormdataBody(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpTokenResponse(rsp)
}

// OpUserinfoWithResponse request returning *OpUserinfoResponse
func (c *ClientWithResponses) OpUserinfoWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*OpUserinfoResponse, error) {
	rsp, err := c.OpUserinfo(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseOpUserinfoResponse(rsp)
}

// RpCallbackWithResponse request returning *RpCallbackResponse
func (c *ClientWithResponses) RpCallbackWithResponse(ctx context.Context, params *RpCallbackParams, reqEditors ...RequestEditorFn) (*RpCallbackResponse, error) {
	rsp, err := c.RpCallback(ctx, params, reqEditors...)
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

// ParseIdpSigninResponse parses an HTTP response from a IdpSigninWithResponse call
func ParseIdpSigninResponse(rsp *http.Response) (*IdpSigninResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &IdpSigninResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseIdpSignupResponse parses an HTTP response from a IdpSignupWithResponse call
func ParseIdpSignupResponse(rsp *http.Response) (*IdpSignupResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &IdpSignupResponse{
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

// ParseOpCallbackResponse parses an HTTP response from a OpCallbackWithResponse call
func ParseOpCallbackResponse(rsp *http.Response) (*OpCallbackResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpCallbackResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseOpLoginResponse parses an HTTP response from a OpLoginWithResponse call
func ParseOpLoginResponse(rsp *http.Response) (*OpLoginResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpLoginResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	return response, nil
}

// ParseOpLoginViewResponse parses an HTTP response from a OpLoginViewWithResponse call
func ParseOpLoginViewResponse(rsp *http.Response) (*OpLoginViewResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpLoginViewResponse{
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

// ParseOpUserinfoResponse parses an HTTP response from a OpUserinfoWithResponse call
func ParseOpUserinfoResponse(rsp *http.Response) (*OpUserinfoResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &OpUserinfoResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest OPUserInfoResponseSchema
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

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

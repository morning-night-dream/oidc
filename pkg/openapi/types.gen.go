// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package openapi

const (
	BearerScopes = "Bearer.Scopes"
)

// Defines values for OpAuthorizeParamsResponseType.
const (
	Code             OpAuthorizeParamsResponseType = "code"
	CodeIdToken      OpAuthorizeParamsResponseType = "code id_token"
	CodeIdTokenToken OpAuthorizeParamsResponseType = "code id_token token"
	CodeToken        OpAuthorizeParamsResponseType = "code token"
	IdToken          OpAuthorizeParamsResponseType = "id_token"
	IdTokenToken     OpAuthorizeParamsResponseType = "id_token token"
	Token            OpAuthorizeParamsResponseType = "token"
)

// Defines values for OpAuthorizeParamsScope.
const (
	Address       OpAuthorizeParamsScope = "address"
	Email         OpAuthorizeParamsScope = "email"
	OfflineAccess OpAuthorizeParamsScope = "offline_access"
	Openid        OpAuthorizeParamsScope = "openid"
	Phone         OpAuthorizeParamsScope = "phone"
	Profile       OpAuthorizeParamsScope = "profile"
)

// Defines values for OpTokenParamsGrantType.
const (
	AuthorizationCode                     OpTokenParamsGrantType = "authorization_code"
	ClientCredentials                     OpTokenParamsGrantType = "client_credentials"
	Password                              OpTokenParamsGrantType = "password"
	RefreshToken                          OpTokenParamsGrantType = "refresh_token"
	UrnIetfParamsOauthGrantTypeDeviceCode OpTokenParamsGrantType = "urn:ietf:params:oauth:grant-type:device_code"
)

// IdPSigninRequestSchema defines model for IdPSigninRequestSchema.
type IdPSigninRequestSchema struct {
	// Password password
	Password string `json:"password"`

	// Username username
	Username string `json:"username"`
}

// IdPSignupRequestSchema defines model for IdPSignupRequestSchema.
type IdPSignupRequestSchema struct {
	// Password password
	Password string `json:"password"`

	// Username username
	Username string `json:"username"`
}

// OPOpenIDConfigurationResponseSchema defines model for OPOpenIDConfigurationResponseSchema.
type OPOpenIDConfigurationResponseSchema struct {
	// AuthorizationEndpoint http://localhost:1234/op/authorize
	AuthorizationEndpoint string `json:"authorization_endpoint"`

	// Issuer http://localhost:1234/op
	Issuer string `json:"issuer"`

	// RevocationEndpoint http://localhost:1234/op/revoke
	RevocationEndpoint string `json:"revocation_endpoint"`

	// TokenEndpoint http://localhost:1234/op/token
	TokenEndpoint string `json:"token_endpoint"`

	// UserinfoEndpoint http://localhost:1234/op/userinfo
	UserinfoEndpoint string `json:"userinfo_endpoint"`
}

// OPTokenResponseSchema https://openid-foundation-japan.github.io/openid-connect-core-1_0.ja.html#TokenResponse
type OPTokenResponseSchema struct {
	// AccessToken access_token
	AccessToken string `json:"access_token"`

	// ExpiresIn expires_in
	ExpiresIn int `json:"expires_in"`

	// IdToken id_token
	IdToken string `json:"id_token"`

	// RefreshToken refresh_token
	RefreshToken string `json:"refresh_token"`

	// TokenType token_type
	TokenType string `json:"token_type"`
}

// OPUserInfoResponseSchema https://openid.net/specs/openid-connect-core-1_0.html#UserInfoResponse
type OPUserInfoResponseSchema struct {
	// Name name
	Name string `json:"name"`

	// Sub sub
	Sub string `json:"sub"`
}

// OpAuthorizeParams defines parameters for OpAuthorize.
type OpAuthorizeParams struct {
	// ResponseType response_type
	ResponseType OpAuthorizeParamsResponseType `form:"response_type" json:"response_type"`

	// Scope scope
	Scope OpAuthorizeParamsScope `form:"scope" json:"scope"`

	// ClientId client_id
	ClientId string `form:"client_id" json:"client_id"`

	// RedirectUri http://localhost:1234/rp/callback
	RedirectUri string `form:"redirect_uri" json:"redirect_uri"`

	// State state
	State *string `form:"state,omitempty" json:"state,omitempty"`
}

// OpAuthorizeParamsResponseType defines parameters for OpAuthorize.
type OpAuthorizeParamsResponseType string

// OpAuthorizeParamsScope defines parameters for OpAuthorize.
type OpAuthorizeParamsScope string

// OpCallbackParams defines parameters for OpCallback.
type OpCallbackParams struct {
	// Id id
	Id string `form:"id" json:"id"`
}

// OpLoginViewParams defines parameters for OpLoginView.
type OpLoginViewParams struct {
	// AuthRequestId auth request id
	AuthRequestId string `form:"auth_request_id" json:"auth_request_id"`
}

// OpTokenParams defines parameters for OpToken.
type OpTokenParams struct {
	// GrantType grant_type
	GrantType OpTokenParamsGrantType `form:"grant_type" json:"grant_type"`

	// Code code
	Code string `form:"code" json:"code"`

	// RedirectUri http://localhost:1234/rp/callback
	RedirectUri string `form:"redirect_uri" json:"redirect_uri"`
}

// OpTokenParamsGrantType defines parameters for OpToken.
type OpTokenParamsGrantType string

// RpCallbackParams defines parameters for RpCallback.
type RpCallbackParams struct {
	// Code code
	Code string `form:"code" json:"code"`

	// State state
	State string `form:"state" json:"state"`
}

// IdpSigninJSONRequestBody defines body for IdpSignin for application/json ContentType.
type IdpSigninJSONRequestBody = IdPSigninRequestSchema

// IdpSignupJSONRequestBody defines body for IdpSignup for application/json ContentType.
type IdpSignupJSONRequestBody = IdPSignupRequestSchema

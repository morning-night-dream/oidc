package op

import (
	"crypto/rsa"

	"github.com/morning-night-dream/oidc/cache"
	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

type OP struct {
	AllowClientID        string
	AllowRedirectURI     string
	AuthorizeParamsCache *cache.Cache[openapi.OpAuthorizeParams]
	UserCache            *cache.Cache[model.User]
	LoggedInUserCache    *cache.Cache[model.User]
	AccessTokenCache     *cache.Cache[model.AccessToken]
	RefreshTokenCache    *cache.Cache[model.RefreshToken]
	IDTokenCache         *cache.Cache[model.IDToken]
	PrivateKey           *rsa.PrivateKey
	Issuer               string
}

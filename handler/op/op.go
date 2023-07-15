package op

import (
	"crypto/rsa"

	"github.com/morning-night-dream/oidc/cache"
	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

type OP struct {
	Issuer               string
	AllowClientID        string
	AllowRedirectURI     string
	SelfURL              string
	PrivateKey           *rsa.PrivateKey
	AuthorizeParamsCache *cache.Cache[openapi.OpAuthorizeParams]
	UserCache            *cache.Cache[model.User] // User情報ストア
	LoggedInUserCache    *cache.Cache[model.User] // ログイン済みのUser情報
	AccessTokenCache     *cache.Cache[model.User] // アクセストークン
	RefreshTokenCache    *cache.Cache[model.User] // リフレッシュトークン
	IDTokenCache         *cache.Cache[model.User] // IDトークン
}

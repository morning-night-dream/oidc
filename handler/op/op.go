package op

import (
	"github.com/morning-night-dream/oidc/cache"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

type OP struct {
	AllowClientID        string
	AllowRedirectURI     string
	AuthorizeParamsCache *cache.Cache[openapi.OpAuthorizeParams]
}

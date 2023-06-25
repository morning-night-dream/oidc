package idp

import (
	"github.com/morning-night-dream/oidc/cache"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

type IdP struct {
	UsernamePasswordCache *cache.Cache[openapi.UsernamePassword]
}

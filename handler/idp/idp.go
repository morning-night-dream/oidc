package idp

import (
	"github.com/morning-night-dream/oidc/cache"
	"github.com/morning-night-dream/oidc/model"
)

type IdP struct {
	UserCache *cache.Cache[model.User]
}

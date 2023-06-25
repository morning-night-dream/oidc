package rp

import (
	"fmt"
	"net/http"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (rp *RP) Callback(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.RpCallbackParams,
) {
	// token取得
	// ここでトークンキャッシュする ???

	// userinfo取得

	w.Write([]byte(fmt.Sprintf("code: %s, state: %s", params.Code, params.State)))
}

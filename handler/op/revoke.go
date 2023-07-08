package op

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/morning-night-dream/oidc/pkg/log"
)

func (op *OP) Revoke(
	w http.ResponseWriter,
	r *http.Request,
) {
	authorization := r.Header.Get("Authorization")

	if authorization == "" {
		log.Log().Warn("authorization header is empty")

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	str := strings.Split(authorization, " ")

	if len(str) != 2 {
		log.Log().Warn(fmt.Sprintf("authorization header is invalid: %s", authorization))

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	if str[0] != "Basic" {
		log.Log().Warn(fmt.Sprintf("authorization header is invalid: %s", authorization))

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	if dec, err := base64.StdEncoding.DecodeString(str[1]); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to decode authorization header: %v", err))

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	} else {
		// TODO client id が適切であるか検証する
		log.Log().Info(fmt.Sprintf("decoded authorization header: %s", dec))
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form:%s", err), http.StatusInternalServerError)

		return
	}

	hint := r.FormValue("token_type_hint")

	token := r.FormValue("token")

	switch hint {
	case "access_token":
		if err := op.AccessTokenCache.Delete(token); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)

			return
		}
	case "refresh_token":
		if err := op.RefreshTokenCache.Delete(token); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)

			return
		}
	default:
		_ = op.AccessTokenCache.Delete(token)
		_ = op.RefreshTokenCache.Delete(token)
	}
}

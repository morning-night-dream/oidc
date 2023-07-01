package op

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/log"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Userinfo(
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

	if str[0] != "Bearer" {
		log.Log().Warn(fmt.Sprintf("authorization header is invalid: %s", authorization))

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	if _, err := model.ParseAccessToken(str[1], "sign"); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to parse access token: %v", err))

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	res := openapi.OPUserInfoResponseSchema{
		Sub:  "sub",
		Name: "name",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to encode response: %v", err))

		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

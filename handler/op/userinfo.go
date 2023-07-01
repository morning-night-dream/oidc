package op

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Userinfo(
	w http.ResponseWriter,
	r *http.Request,
) {
	authorization := r.Header.Get("Authorization")

	if authorization == "" {
		log.Printf("authorization header is empty")

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	str := strings.Split(authorization, " ")

	if len(str) != 2 {
		log.Printf("authorization header is invalid: %s", authorization)

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	if str[0] != "Bearer" {
		log.Printf("authorization header is invalid: %s", authorization)

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	if _, err := model.ParseAccessToken(str[1], "sign"); err != nil {
		log.Printf("failed to parse access token: %v", err)

		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	res := openapi.OPUserInfoResponseSchema{
		Sub:  "sub",
		Name: "name",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("failed to encode response: %v", err)

		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

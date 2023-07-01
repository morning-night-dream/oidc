package op

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Token(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpTokenParams,
) {
	log.Printf("%+v", params)

	// ref. https://qiita.com/TakahikoKawasaki/items/970548727761f9e02bcd
	// 1.3 hybrid type で実装してみる
	// -> アクセストークンを revoke したいため
	at := model.GenerateAccessToken(
		"iss",
		"sub",
		"aud",
		"jti",
		"scope",
		"client_id",
	)

	rt := model.GenerateRefreshToken()

	it := model.GenerateIDToken(
		"iss",
		"sub",
		"aud",
		"nonce",
		"name",
	)

	res := openapi.OPTokenResponseSchema{
		TokenType:    "Bearer",
		AccessToken:  at.JWT("sign"),
		IdToken:      it.JWT("sign"),
		RefreshToken: rt.Base64(),
		ExpiresIn:    3600,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("failed to encode response: %v", err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

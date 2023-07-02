package op

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/morning-night-dream/oidc/model"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Token(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form:%s", err), http.StatusInternalServerError)

		return
	}

	code := r.FormValue("code")

	user, err := op.LoggedInUserCache.Get(code)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	authReq, err := op.AuthorizeParamsCache.Get(code)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
	}

	// ref. https://qiita.com/TakahikoKawasaki/items/970548727761f9e02bcd
	// 1.3 hybrid type で実装してみる
	// -> アクセストークンを revoke したいため
	at := model.GenerateAccessToken(
		op.Issuer,
		user.ID,
		op.AllowClientID,
		"jti",
		"scope",
		"client_id",
	)

	if err := op.AccessTokenCache.Set(user.ID, at); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	rt := model.GenerateRefreshToken()

	if err := op.RefreshTokenCache.Set(user.ID, rt); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	it := model.GenerateIDToken(
		op.Issuer,
		user.ID,
		op.AllowClientID,
		*authReq.Nonce,
		user.Username,
	)

	if err := op.IDTokenCache.Set(user.ID, it); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

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

		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

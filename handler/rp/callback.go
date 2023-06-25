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
	tRes, err := http.Get(rp.TokenURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer tRes.Body.Close()

	// tBody, _ := io.ReadAll(tRes.Body)

	// var token openapi.OpTokenResponse
	// if err := json.Unmarshal(tBody, &token); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)

	// 	return
	// }

	// userinfo取得
	uRes, err := http.Get(rp.UserInfoURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer uRes.Body.Close()

	// uBody, _ := io.ReadAll(uRes.Body)

	// var userinfo openapi.OpUserinfoResponse
	// if err := json.Unmarshal(uBody, &userinfo); err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)

	// 	return
	// }

	w.Write([]byte(fmt.Sprintf("%+v", "userinfo")))
}

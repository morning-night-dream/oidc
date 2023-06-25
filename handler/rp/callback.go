package rp

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (rp *RP) Callback(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.RpCallbackParams,
) {
	var buf bytes.Buffer

	buf.WriteString(rp.TokenURL)

	values := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {params.Code},
		"redirect_uri": {rp.RedirectURI},
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	url := buf.String()

	tRes, err := http.Get(url)
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

package rp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/morning-night-dream/oidc/pkg/log"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (rp *RP) Callback(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.RpCallbackParams,
) {
	values := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {params.Code},
		"redirect_uri": {rp.RedirectURI},
	}

	tRes, err := http.Post(rp.TokenURL, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		log.Log().Warn(fmt.Sprintf("failed to request: %v", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer tRes.Body.Close()

	tBody, _ := io.ReadAll(tRes.Body)

	var token openapi.OPTokenResponseSchema
	if err := json.Unmarshal(tBody, &token); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to unmarshal: %v", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	log.Log().Debug(fmt.Sprintf("%+v", token))

	// userinfo取得
	client := &http.Client{
		Transport: NewAuthorizationTransport(token.AccessToken),
	}

	req, err := http.NewRequest(http.MethodGet, rp.UserInfoURL, nil)
	if err != nil {
		log.Log().Warn(fmt.Sprintf("failed to create request: %v", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	res, err := client.Do(req)
	if err != nil {
		log.Log().Warn(fmt.Sprintf("failed to request: %v", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var userinfo openapi.OPUserInfoResponseSchema
	if err := json.Unmarshal(body, &userinfo); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to unmarshal: %v", err))

		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte(fmt.Sprintf("%+v", userinfo)))
}

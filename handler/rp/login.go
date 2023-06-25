package rp

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (rp *RP) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	var buf bytes.Buffer

	buf.WriteString(rp.AuthURL)

	state := uuid.NewString()

	http.SetCookie(w, &http.Cookie{
		Name:     "state",
		Value:    state,
		Path:     "/",
		Domain:   "localhost",
		Expires:  time.Now().Add(time.Hour),
		Secure:   true,
		HttpOnly: true,
	})

	values := url.Values{
		"response_type": {"code"},                       // Authorization Flow なので code を指定
		"client_id":     {rp.ClientID},                  // RPを識別するためのID OPに登録してある必要がある
		"redirect_uri":  {rp.RedirectURI},               // ログイン後にリダイレクトさせるURL OPに登録してある必要がある
		"scope":         {strings.Join(rp.Scopes, " ")}, // RPが要求するスコープ OPに登録してある必要がある
		"state":         {state},                        // CSRF対策のためのstate
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	url := buf.String()

	http.Redirect(w, r, url, http.StatusFound)
}

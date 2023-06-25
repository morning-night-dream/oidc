package op

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func (op *OP) Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form:%s", err), http.StatusInternalServerError)
		return
	}

	id := r.FormValue("id")

	username := r.FormValue("username")

	password := r.FormValue("password")

	log.Printf("%s, %s, %s", id, username, password)

	// TODO: usernameとpasswordの検証

	// id でログインに成功していることをキャッシュに保存しておく

	var buf bytes.Buffer

	buf.WriteString("http://localhost:1234/op/callback")

	values := url.Values{
		"id": {id},
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	url := buf.String()

	http.Redirect(w, r, url, http.StatusFound)
}

package op

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/morning-night-dream/oidc/model"
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

	// usernameとpasswordの検証
	user, err := op.UserCache.Get(username)
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	if err := model.CompareHashAndPassword(user.Password, password); err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)

		return
	}

	// user がログインに成功していることをキャッシュに保存しておく
	if err := op.LoggedInUserCache.Set(id, user); err != nil {
		http.Error(w, fmt.Sprintf("cannot set cache:%s", err), http.StatusInternalServerError)

		return
	}

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s/op/callback", op.SelfURL))

	values := url.Values{
		"id": {id},
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	url := buf.String()

	http.Redirect(w, r, url, http.StatusFound)
}

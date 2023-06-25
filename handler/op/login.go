package op

import (
	"fmt"
	"log"
	"net/http"
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
}

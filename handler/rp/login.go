package rp

import "net/http"

func Login(
	w http.ResponseWriter,
	r *http.Request,
) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

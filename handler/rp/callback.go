package rp

import "net/http"

func (rp *RP) Callback(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Write([]byte("callback"))
}

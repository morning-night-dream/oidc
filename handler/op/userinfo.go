package op

import "net/http"

func (op *OP) Userinfo(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Write([]byte("userinfo"))
}

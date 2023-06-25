package op

import (
	"net/http"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Token(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpTokenParams,
) {
	w.Write([]byte("token"))
}

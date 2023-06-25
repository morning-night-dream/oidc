package op

import (
	"log"
	"net/http"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Authorize(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpAuthorizeParams,
) {
	log.Printf("%+v", params)

	// Client ID の検証

	// Redirect URL の検証

	log.Printf("%+v", r.URL.Query())
}

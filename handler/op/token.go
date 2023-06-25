package op

import (
	"log"
	"net/http"
)

func (op *OP) Token(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Client ID の検証

	// Redirect URL の検証

	log.Printf("%+v", r.URL.Query())
}

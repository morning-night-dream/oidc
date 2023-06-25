package op

import (
	"log"
	"net/http"
)

func Authorize(
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Printf("%+v", r.URL.Query())
}

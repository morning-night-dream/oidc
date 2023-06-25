package op

import (
	"log"
	"net/http"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) LoginView(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpLoginViewParams,
) {
	// ログイン画面を表示する

	log.Printf("%+v", params)

	w.Write([]byte(params.AuthRequestId))
}

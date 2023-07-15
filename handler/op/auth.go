package op

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Authorize(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpAuthorizeParams,
) {
	// TODO: Client ID の検証

	// TODO: Redirect URL の検証

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%s/op/login", op.SelfURL))

	id := uuid.NewString()

	op.AuthorizeParamsCache.Set(id, params)

	values := url.Values{
		"auth_request_id": {id},
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	url := buf.String()

	http.Redirect(w, r, url, http.StatusFound)
}

package op

import (
	"bytes"
	"net/http"
	"net/url"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Callback(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpCallbackParams,
) {
	// request id に紐づく auth request を取得
	authReq, _ := op.AuthorizeParamsCache.Get(params.Id)

	// TODO: このセッション(リクエスト)が認証されたかどうかを判定する

	// TODO: response_type=token の場合は access_token を返す

	var buf bytes.Buffer

	buf.WriteString(authReq.RedirectUri)

	values := url.Values{
		"code":  {"code"},  // これはなに?
		"state": {"state"}, // これはなに?
	}

	buf.WriteByte('?')

	buf.WriteString(values.Encode())

	url := buf.String()

	http.Redirect(w, r, url, http.StatusFound)
}

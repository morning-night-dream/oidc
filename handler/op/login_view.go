package op

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) LoginView(
	w http.ResponseWriter,
	r *http.Request,
	params openapi.OpLoginViewParams,
) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot parse form:%s", err), http.StatusInternalServerError)
		return
	}

	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	data := &struct {
		ID    string
		Error string
	}{
		ID:    params.AuthRequestId,
		Error: errMsg,
	}
	if err = loginTmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

const tmp = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Login</title>
	</head>
	<body style="display: flex; align-items: center; justify-content: center; height: 100vh;">
		<form method="POST" action="/login" style="height: 200px; width: 200px;">

			<input type="hidden" name="id" value="{{.ID}}">

			<div>
				<label for="username">Username:</label>
				<input id="username" name="username" style="width: 100%">
			</div>

			<div>
				<label for="password">Password:</label>
				<input id="password" name="password" style="width: 100%">
			</div>

			<p style="color:red; min-height: 1rem;">{{.Error}}</p>

			<button type="submit">Login</button>
		</form>
	</body>
</html>
`

var loginTmpl, _ = template.New("login").Parse(tmp)

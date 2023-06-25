package op

import (
	"log"
	"net/http"
)

// LoginView is a view for login page.
func LoginView(
	w http.ResponseWriter,
	r *http.Request,
) {
	log.Printf("%+v", r.URL.Query())

	w.Write([]byte(`
		<html>
			<head>
				<title>Login In</title>
			</head>
			<body>
			    <!--ここのリクエストパラメータに何かを載っける必要がありそう-->
				<form action="/op/login" method="post">
					<input type="text" name="username" />
					<input type="password" name="password" />	
					<input type="submit" value="Login" />
				</form>
			</body>
		</html>
	`))
}

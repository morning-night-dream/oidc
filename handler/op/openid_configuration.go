package op

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/morning-night-dream/oidc/pkg/log"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) OpenIDConfiguration(
	w http.ResponseWriter,
	r *http.Request,
) {
	res := openapi.OPOpenIDConfigurationResponseSchema{
		AuthorizationEndpoint: fmt.Sprintf("%s/op/authorize", op.SelfURL),
		Issuer:                op.SelfURL,
		JwksUrl:               fmt.Sprintf("%s/op/jwks", op.SelfURL),
		RevocationEndpoint:    fmt.Sprintf("%s/op/revoke", op.SelfURL),
		TokenEndpoint:         fmt.Sprintf("%s/op/token", op.SelfURL),
		UserinfoEndpoint:      fmt.Sprintf("%s/op/userinfo", op.SelfURL),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to encode response: %v", err))

		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

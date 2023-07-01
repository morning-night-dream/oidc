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
		Issuer:                "http://localhost:1234",
		AuthorizationEndpoint: "http://localhost:1234/op/authorize",
		TokenEndpoint:         "http://localhost:1234/op/token",
		UserinfoEndpoint:      "http://localhost:1234/op/userinfo",
		RevocationEndpoint:    "http://localhost:1234/op/revoke",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to encode response: %v", err))

		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

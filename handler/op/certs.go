package op

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/morning-night-dream/oidc/pkg/log"
	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func (op *OP) Certs(
	w http.ResponseWriter,
	r *http.Request,
) {
	pub := op.PrivateKey.PublicKey

	res := openapi.OPJWKSetResponseSchema{
		Keys: []openapi.OPJWKSetKey{
			{
				Alg: "RS256",
				E:   EncodeUint64ToString(uint64(pub.E)),
				Kid: "12345678",
				Kty: "RSA",
				N:   EncodeToString(pub.N.Bytes()),
				Use: "sig",
			},
		},
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Log().Warn(fmt.Sprintf("failed to encode response: %v", err))

		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func EncodeUint64ToString(v uint64) string {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, v)

	i := 0
	for ; i < len(data); i++ {
		if data[i] != 0x0 {
			break
		}
	}

	return EncodeToString(data[i:])
}

func EncodeToString(src []byte) string {
	return base64.RawURLEncoding.EncodeToString(src)
}

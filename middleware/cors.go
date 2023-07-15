package middleware

import (
	"github.com/go-chi/cors"

	"github.com/morning-night-dream/oidc/pkg/openapi"
)

func NewCORS(
	selfURL string,
) openapi.MiddlewareFunc {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{selfURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            false,
		MaxAge:           300,
	})
}

package middleware

import (
	c "book-store-api/api/constants"
	"book-store-api/api/utils"
	"book-store-api/handler"
	"net/http"
)

const (
	ApiKeyHeader = "ApiKey"
)

func ApiKeyMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get(ApiKeyHeader)
		if apiKey == "" {
			handler.HttpError(w, http.StatusForbidden, c.UNAUTHRIZED_REQUEST, r.URL)
			return
		}
		_, err := utils.ValidateEnvVar(ApiKeyHeader, apiKey)
		if err != nil {
			handler.HttpError(w, http.StatusForbidden, c.UNAUTHRIZED_REQUEST, r.URL)
			return
		}
		next.ServeHTTP(w, r)
	})
}

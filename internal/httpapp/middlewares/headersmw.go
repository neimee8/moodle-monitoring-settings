package middlewares

import (
	"net/http"
	"strings"
)

type HeadersMw struct{}

func NewHeadersMw() HeadersMw {
	return HeadersMw{}
}

func (mw HeadersMw) Mw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlParts := strings.Split(r.URL.Path, "/")

		if len(urlParts) > 1 && urlParts[1] == "api" {
			w.Header().Set("Content-Type", "application/json")
		}

		next.ServeHTTP(w, r)
	})
}

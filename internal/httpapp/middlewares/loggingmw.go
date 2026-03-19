package middlewares

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"settings/internal/config"
	"settings/internal/httpapp/wrappers"
	"time"

	"github.com/google/uuid"
)

type LoggingMw struct {
	cfg *config.Config
}

func NewLoggingMw(cfg *config.Config) LoggingMw {
	return LoggingMw{
		cfg: cfg,
	}
}

func (mw LoggingMw) Mw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body []byte

		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)

			if err != nil {
				http.Error(w, "failed to read request body", http.StatusBadRequest)
				return
			}

			body = bodyBytes
			r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		}

		uuid := uuid.New()

		rw := wrappers.NewResponseWriter(w)

		fmt.Printf(
			"🕰️ [%s]\n⬇️  Got request:\n📍 request uuid: %s\n⚙️  method: %s\n🗺️  path: %s\n💿  body:\n%s\n\n",
			time.Now().Local().Format(mw.cfg.TimeFormat),
			uuid,
			r.Method,
			r.URL.Path,
			string(body),
		)

		next.ServeHTTP(rw, r)

		fmt.Printf(
			"🕰️ [%s]\n⬆️  Prepared response:\n📍 request uuid: %s\nℹ️  status: %d\n💿 body:\n%s\n\n",
			time.Now().Local().Format(mw.cfg.TimeFormat),
			uuid,
			rw.Status,
			rw.Body.String(),
		)
	})
}

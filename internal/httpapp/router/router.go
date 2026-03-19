package router

import (
	"net/http"
	"settings/internal/httpapp/handlers"
	"settings/internal/httpapp/middlewares"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h handlers.Handlers, m middlewares.Middlewares) *chi.Mux {
	r := chi.NewRouter()

	mws(r, m)
	routes(r, h)

	return r
}

func routes(
	r *chi.Mux,
	h handlers.Handlers,
) {
	routesApi(r, h.Api)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		resp := []byte("{\"msg\": \"not found\", \"data\": \"\"}")
		w.Write(resp)
	})

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		resp := []byte("{\"msg\": \"method not allowed\", \"data\": \"\"}")
		w.Write(resp)
	})
}

func mws(
	r *chi.Mux,
	mws middlewares.Middlewares,
) {
	r.Use(mws.Logging.Mw)
	r.Use(mws.Response.Mw)
	r.Use(mws.Backup.Mw)
	r.Use(mws.Headers.Mw)
}

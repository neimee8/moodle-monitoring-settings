package router

import (
	"net/http"
	"settings/internal/httpapp/handlers/api"

	"github.com/go-chi/chi/v5"
)

func routesApi(
	r *chi.Mux,
	h api.ApiHandlers,
) {
	r.Handle("/api/settings", http.HandlerFunc(h.HandleSettings))
	r.Get("/api/rollback", http.HandlerFunc(h.HandleRollback))
}

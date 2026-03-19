package handlers

import (
	"settings/internal/config"
	"settings/internal/engine"
	"settings/internal/httpapp/handlers/api"
)

type Handlers struct {
	Api api.ApiHandlers
}

func NewHandlers(cfg *config.Config, cmdCh chan engine.Cmd) Handlers {
	return Handlers{
		Api: api.NewApiHandlers(cfg, cmdCh),
	}
}

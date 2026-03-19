package api

import (
	"settings/internal/config"
	"settings/internal/engine"
)

type ApiHandlers struct {
	cfg   *config.Config
	cmdCh chan engine.Cmd
}

func NewApiHandlers(cfg *config.Config, cmdCh chan engine.Cmd) ApiHandlers {
	return ApiHandlers{
		cfg:   cfg,
		cmdCh: cmdCh,
	}
}

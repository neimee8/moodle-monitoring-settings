package middlewares

import (
	"settings/internal/config"
	"settings/internal/engine"
)

type Middlewares struct {
	Backup   BackupMw
	Logging  LoggingMw
	Headers  HeadersMw
	Response ResponseMw
}

func NewMiddlewares(cfg *config.Config, cmdCh chan engine.Cmd) Middlewares {
	return Middlewares{
		Backup:   NewBackupMw(cfg, cmdCh),
		Logging:  NewLoggingMw(cfg),
		Headers:  NewHeadersMw(),
		Response: NewResponseMw(cfg),
	}
}

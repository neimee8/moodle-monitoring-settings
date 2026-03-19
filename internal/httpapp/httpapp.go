package httpapp

import (
	"context"
	"errors"
	"net/http"
	"settings/internal/config"
	"settings/internal/engine"
	"settings/internal/httpapp/handlers"
	"settings/internal/httpapp/middlewares"
	"settings/internal/httpapp/router"
)

type HttpApp struct {
	server *http.Server
}

func New(cfg *config.Config, cmdCh chan engine.Cmd) *HttpApp {
	h := handlers.NewHandlers(cfg, cmdCh)
	mws := middlewares.NewMiddlewares(cfg, cmdCh)
	r := router.NewRouter(h, mws)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: r,
	}

	return &HttpApp{
		server: srv,
	}
}

func (a *HttpApp) Start() error {
	err := a.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

func (a *HttpApp) Shutdown(ctx context.Context) error {
	return a.server.Shutdown(ctx)
}

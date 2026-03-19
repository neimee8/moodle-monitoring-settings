package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"settings/internal/config"
	"settings/internal/engine"
	"settings/internal/httpapp"
	"sync"
	"syscall"
	"time"
)

func main() {
	cfg := config.Load()

	var workerWG sync.WaitGroup
	workerWG.Add(1)
	cmdCh := make(chan engine.Cmd, 128)

	go engine.Worker(cmdCh, &workerWG)

	app := httpapp.New(cfg, cmdCh)
	serverErrCh := make(chan error, 1)

	go func() {
		serverErrCh <- app.Start()
	}()

	fmt.Printf("🕰️ [%s]\n🚀 Server started\n\n", time.Now().Local().Format(cfg.TimeFormat))

	sigCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case <-sigCtx.Done():
		fmt.Printf("🕰️ [%s]\n🔽 Shutdown signal received\n\n", time.Now().Local().Format(cfg.TimeFormat))

	case err := <-serverErrCh:
		if err != nil {
			fmt.Printf("🕰️ [%s] ❌\nHTTP server stopped with error: %s\n\n", time.Now().Local().Format(cfg.TimeFormat), err.Error())
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Shutdown(shutdownCtx); err != nil {
		fmt.Printf("🕰️ [%s]\n❌ Graceful http shutdown failed: %s\n\n", time.Now().Local().Format(cfg.TimeFormat), err.Error())
	}

	close(cmdCh)
	workerWG.Wait()

	fmt.Printf("🕰️ [%s]\n🛑 Graceful shutdown complete\n\n", time.Now().Local().Format(cfg.TimeFormat))
}

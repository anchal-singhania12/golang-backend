package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"gitlab.com/fanligafc-group/fanligafc-backend/internal/config"
)

type App struct {
	config *config.Config
	server *http.Server
}

func NewApp() (*App, error) {

	// Load dependencies
	deps, err := InitializeDeps()
	if err != nil {
		log.Fatalf("error in initializing dependencies: %v", err)
		return nil, err
	}

	handler := InitiateHandlers(deps)
	routes := InitializeRoutesEngine(handler)

	server := &http.Server{
		Addr:    ":" + deps.cfg.HTTP.Addr,
		Handler: routes,
	}

	return &App{
		config: deps.cfg,
		server: server,
	}, nil
}

func (app *App) Run() error {
	// Start server in background
	go func() {
		log.Printf("🚀 starting server on %s\n", app.config.HTTP.Addr)
		if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for SIGINT or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("🔌 shutting down server…")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.server.Shutdown(ctx); err != nil {
		log.Printf("forced to shutdown: %v", err)
		return err
	}
	log.Println("👋 server stopped cleanly")
	return nil
}

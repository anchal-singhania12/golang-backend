package main

import (
	"log"

	"gitlab.com/fanligafc-group/fanligafc-backend/internal/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("error during shutdown: %v", err)
	}
}

package app

import (
	"log"

	"gitlab.com/fanligafc-group/fanligafc-backend/internal/config"
	"gitlab.com/fanligafc-group/fanligafc-backend/pkg/database"
	tokenmanager "gitlab.com/fanligafc-group/fanligafc-backend/pkg/token_manager"
	"gorm.io/gorm"
)

type Dependencies struct {
	cfg                *config.Config
	db                 *gorm.DB
	accessTokenManager tokenmanager.TokenManager
}

func InitializeDeps() (*Dependencies, error) {
	cfg := config.LoadConfig()

	// Initialize database connection
	gormDB, err := database.InitializeDB(cfg)
	if err != nil {
		log.Fatalf("Error in intializing DB: %v", err)
		return nil, err
	}
	log.Printf("Database Initialized")
	return &Dependencies{
		cfg: cfg,
		db:  gormDB,
		//accessTokenManager: accessTokenManager,
	}, nil
}

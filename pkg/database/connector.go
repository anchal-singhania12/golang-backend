package database

import (
	"fmt"

	"gitlab.com/fanligafc-group/fanligafc-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB(cfg *config.Config) (*gorm.DB, error) {

	//migrate db up
	err := MigrateDB(&cfg.Database)
	if err != nil {
		return nil, err
	}

	//initialize db instance
	gormDB, err := NewGORM(&cfg.Database)
	if err != nil {
		return nil, err
	}

	return gormDB, nil
}

func NewGORM(cfg *config.Database) (*gorm.DB, error) {
	dsn := cfg.URL
	if dsn == "" {
		// fallback to building DSN
		dsn = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.Driver, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
		)
	}
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

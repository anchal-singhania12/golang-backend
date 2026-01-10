package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gitlab.com/fanligafc-group/fanligafc-backend/internal/config"
)

func MigrateDB(dbCfg *config.Database) error {
	sqlDB, err := sql.Open(dbCfg.Driver, dbCfg.URL)
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("postgres driver: %w", err)
	}

	path := dbCfg.MigrationsPath
	if path == "" {
		path = "pkg/database/migrations"
	}
	if len(path) < 7 || path[:7] != "file://" {
		path = "file://" + path
	}

	m, err := migrate.NewWithDatabaseInstance(path, dbCfg.Driver, driver)
	if err != nil {
		return fmt.Errorf("migrator init: %w", err)
	}

	if err := MigrateUp(m); err != nil {
		return err
	}

	return nil
}

func MigrateUp(m *migrate.Migrate) error {
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

func MigrateDown(m *migrate.Migrate) error {
	if err := m.Down(); err != nil {
		return fmt.Errorf("migrate down: %w", err)
	}
	return nil
}

func MigrateSteps(m *migrate.Migrate, n int) error {
	if err := m.Steps(n); err != nil {
		return fmt.Errorf("migrate steps: %w", err)
	}
	return nil
}

package repository

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/nekvil/cars-go-api/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const migrationsPath = "file://migrations/"

func SetupDatabase() *gorm.DB {
	utils.LoadEnv()
	dsn := getDSN()

	utils.Logger.Info("Setting up database...")
	if err := applyMigrations(dsn); err != nil {
		utils.Logger.Fatalf("Failed to apply migrations: %v", err)
	}

	var db *gorm.DB
	var err error

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		utils.Logger.Fatalf("Failed to connect to database: %v", err)
	}

	utils.Logger.Info("Database setup completed")

	return db
}

func getDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		utils.GetEnv("DB_USER"),
		utils.GetEnv("DB_PASSWORD"),
		utils.GetEnv("DB_HOST"),
		utils.GetEnv("DB_PORT"),
		utils.GetEnv("DB_NAME"),
		utils.GetEnv("DB_SSLMODE"))
}

func applyMigrations(dsn string) error {
	utils.Logger.Info("Applying database migrations")

	m, err := migrate.New(migrationsPath, dsn)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	utils.Logger.Info("Database migrations applied successfully")

	return nil
}

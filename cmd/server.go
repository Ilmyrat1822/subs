package cmd

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/Ilmyrat1822/subs/internal/config"
	"github.com/Ilmyrat1822/subs/internal/database"
	"github.com/Ilmyrat1822/subs/utils/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Echo     *echo.Echo
	Config   *config.Schema
	Database *gorm.DB
}

func NewServer() *Server {
	cfg := config.GetConfig()

	// Ensure DB exists before connecting
	if err := EnsureDBExists(cfg.PostgresUri); err != nil {
		log.Fatalf("failed to ensure database exists: %v", err)
	}

	// Connect to the database
	db, err := database.Connect(cfg.PostgresUri)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	e := echo.New()
	e.Validator = validator.NewValidator()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	return &Server{
		Echo:     e,
		Config:   cfg,
		Database: db,
	}
}

func EnsureDBExists(dsn string) error {
	parsed, err := url.Parse(dsn)
	if err != nil {
		return err
	}

	dbName := strings.TrimPrefix(parsed.Path, "/")
	if dbName == "" {
		return fmt.Errorf("database name is empty in DSN")
	}

	parsed.Path = "/postgres"
	tmpDsn := parsed.String()

	db, err := gorm.Open(postgres.Open(tmpDsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to postgres DB: %w", err)
	}

	// Create database if not exists
	err = db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\";", dbName)).Error
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 42P04") {
			return nil
		}
		return fmt.Errorf("failed to create database: %w", err)
	}

	log.Printf("Database %s exists or created successfully", dbName)
	return nil
}

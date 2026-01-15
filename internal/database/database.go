package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger: gormLogger.Default.LogMode(gormLogger.Info),
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			PrepareStmt:     false,
			CreateBatchSize: 100,
		},
	)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(40)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	// ONLY extensions, no schema
	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL")
	return db, nil
}

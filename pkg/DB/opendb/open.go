package opendb

import (
	"fmt"
	"log"
	password "password-manager/internal/Password"
	"password-manager/pkg/DB/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() error {
	cfg := config.Load()

	connStr := cfg.GetConnectionString()

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}

	var err error
	DB, err = gorm.Open(postgres.Open(connStr), config)
	if err != nil {
		return fmt.Errorf("failed to connectiom to database: %v", err)
	}

	log.Println("Succsesfully to connection to database ✅")

	DB.Exec("CREATE SCHEMA IF NOT EXISTS pass_manager")
	DB.Exec("SET search_path TO pass_manager")

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = DB.AutoMigrate(&password.Password{})
	if err != nil {
		return fmt.Errorf("failed AutoMigrate: %v", err)
	}

	log.Println("Database migrated successfuly ✅")
	return nil
}

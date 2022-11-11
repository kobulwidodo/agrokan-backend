package sql

import (
	"agrokan-backend/src/business/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
	SSLMode  string
}

func Init(cfg Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai", cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal(err.Error())
	}

	return db
}

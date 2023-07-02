package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/andresianipar/go-template/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	port, err := strconv.ParseInt(configs.Get("DB_PORT"), 10, 32)
	if err != nil {
		log.Fatal("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		configs.Get("DB_HOST"),
		configs.Get("DB_USER"),
		configs.Get("DB_PASSWORD"),
		configs.Get("DB_NAME"),
		port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	return db
}

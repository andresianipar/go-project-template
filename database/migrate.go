package database

import (
	"log"

	"github.com/andresianipar/go-template/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if !db.Migrator().HasTable(&model.Migration{}) {
		if err := db.Migrator().CreateTable(&model.Migration{}); err != nil {
			log.Fatalf("database migration failed: %s", err)
		}
	}
}

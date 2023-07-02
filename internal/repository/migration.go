package repository

import (
	"github.com/andresianipar/go-template/internal/model"
	"gorm.io/gorm"
)

type MigrationRepository interface {
	GetMigrations() (*[]model.Migration, error)
}

type migrationRepository struct {
	db *gorm.DB
}

func NewMigrationRepository(db *gorm.DB) MigrationRepository {
	return &migrationRepository{db: db}
}

func (mr *migrationRepository) GetMigrations() (*[]model.Migration, error) {
	var migrations []model.Migration
	mr.db.Find(&migrations)
	return &migrations, nil
}

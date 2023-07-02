package repository

import "gorm.io/gorm"

type Repository struct {
	migrationRepository MigrationRepository
}

func New(db *gorm.DB) *Repository {
	return &Repository{
		NewMigrationRepository(db),
	}
}

func (r *Repository) MigrationRepository() MigrationRepository {
	return r.migrationRepository
}

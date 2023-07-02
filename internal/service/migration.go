package service

import (
	"github.com/andresianipar/go-template/internal/model"
	"github.com/andresianipar/go-template/internal/repository"
)

type MigrationService interface {
	RetrieveMigrations() (*[]model.Migration, error)
}

type migrationService struct {
	repository *repository.Repository
}

func NewMigrationService(repository *repository.Repository) MigrationService {
	return &migrationService{
		repository: repository,
	}
}

func (ms *migrationService) RetrieveMigrations() (*[]model.Migration, error) {
	migrations, err := ms.repository.MigrationRepository().GetMigrations()
	if err != nil {
		return nil, err
	}

	return migrations, nil
}

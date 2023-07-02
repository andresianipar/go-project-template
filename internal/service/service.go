package service

import "github.com/andresianipar/go-template/internal/repository"

type Service struct {
	migrationService MigrationService
}

func New(repository *repository.Repository) *Service {
	return &Service{
		NewMigrationService(repository),
	}
}

func (s *Service) MigrationService() MigrationService {
	return s.migrationService
}

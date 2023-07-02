package view

import (
	"time"

	"github.com/andresianipar/go-template/internal/model"
	"github.com/gofiber/fiber/v2"
)

type Migration struct {
	ID        uint      `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	Name      string    `json:"name"`
}

func (v *View) RetrieveMigrationsResponse(migrations *[]model.Migration) *fiber.Map {
	var serializedMigrations []Migration
	for _, migration := range *migrations {
		serializedMigrations = append(serializedMigrations, Migration{
			ID:        migration.ID,
			Timestamp: migration.Timestamp,
			Name:      migration.Name,
		})
	}

	return &fiber.Map{
		"data":   &serializedMigrations,
		"error":  nil,
		"status": true,
	}
}

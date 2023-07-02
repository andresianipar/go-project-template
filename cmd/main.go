package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	_ "github.com/andresianipar/go-template/api"
	"github.com/andresianipar/go-template/configs"
	"github.com/andresianipar/go-template/database"
	app "github.com/andresianipar/go-template/internal"
	"github.com/andresianipar/go-template/internal/repository"
	"github.com/andresianipar/go-template/internal/router"
	"github.com/andresianipar/go-template/internal/service"
	"github.com/andresianipar/go-template/internal/view"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("failed to load .env file")
		}
	}

	// Open a database connection and migrate the schema
	db := database.Connect()
	database.Migrate(db)

	// Go Struct and Field validators
	validator := validator.New()

	// Bootstrap the application
	app := app.New(
		fiber.New(),
		router.New(
			service.New(
				repository.New(db),
			),
			view.New(validator),
		),
	)

	// Serves HTTP requests from a different goroutine
	go func() {
		port, err := strconv.ParseInt(configs.Get("APP_PORT"), 10, 32)
		if err != nil {
			log.Fatal("failed to parse application port")
		}
		if err := app.Server().Listen(fmt.Sprintf(":%d", port)); err != nil {
			log.Panic(err)
		}
	}()

	// Graceful shutdown
	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt, syscall.SIGTERM)

	<-osSignalChan
	_ = app.Server().Shutdown()
	if DB, err := db.DB(); err == nil {
		_ = DB.Close()
	}
	log.Println("App shuts down gracefully")
}

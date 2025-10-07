package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Buckshot-77/expertscrud/src/handlers"
	"github.com/Buckshot-77/expertscrud/src/repositories"
	"github.com/Buckshot-77/expertscrud/src/services"
	"github.com/Buckshot-77/expertscrud/src/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnvOrDefault(key, fallback string) string {
	value, success := os.LookupEnv(key)
	if !success {
		return fallback
	}
	return value
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	server := fiber.New()
	serverPort := getEnvOrDefault("EXPERTSCRUD_APPLICATION_PORT", "3000")

	dialector := postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("EXPERTSCRUD_DATABASE_HOST"),
		os.Getenv("EXPERTSCRUD_DATABASE_USER"),
		os.Getenv("EXPERTSCRUD_DATABASE_PASSWORD"),
		os.Getenv("EXPERTSCRUD_DATABASE_NAME"),
		os.Getenv("EXPERTSCRUD_DATABASE_PORT"),
	))

	db, err := gorm.Open(dialector)

	if err != nil {
		log.Fatalf("error while connecting to the database. Err=%v", err.Error())
	}

	db.AutoMigrate(&structs.Installment{})

	repositoryContainer := repositories.GetRepositories(db)
	servicesContainer := services.GetServices(repositoryContainer)

	handlers.NewHandlerContainer(server, servicesContainer)

	server.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON("Pong!")
	})

	if err := server.Listen(fmt.Sprintf(":%v", serverPort)); err != nil {
		log.Fatalf("error while initializing the server. Err=%v", err.Error())
	}

}

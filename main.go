package main

import (
	"log"

	"HR-management/database"
	"HR-management/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	mongoDb, err := database.InitDatabase();

	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	routes.EmployeeRoutes(app, mongoDb)

	log.Fatal(app.Listen(":8080"))
}
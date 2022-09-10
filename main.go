package main

import (
	"github.com/gofiber/fiber/v2"
	"go-server-jwt/database"
	"go-server-jwt/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8080")
}

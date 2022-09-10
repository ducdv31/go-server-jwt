package main

import "github.com/gofiber/fiber/v2"
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

func main() {
	_, err := gorm.Open(mysql.Open("root:DucDV@31@/mysql"), &gorm.Config{})

	if err != nil {
		panic(any("Could not connect to the database"))
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8080")
}

package main

import (
	"log"

	"bitbucket.org/nithursika/go-fiber-tutorial/database"
	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API")
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	app.Get("/api", Welcome)
	log.Fatal(app.Listen(":3000"))
}

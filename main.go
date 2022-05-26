package main

import (
	"github.com/galifornia/go-simple-crm/database"
	"github.com/galifornia/go-simple-crm/lead"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	lead.SetupLeadRoutes(app)
}

func main() {
	database.OpenDB()
	app := fiber.New()

	setupRoutes(app)

	app.Listen("localhost:3000")
}

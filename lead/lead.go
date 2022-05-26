package lead

import (
	"github.com/galifornia/go-simple-crm/database"
	"github.com/galifornia/go-simple-crm/types"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SetupLeadRoutes(app *fiber.App) {
	leads := app.Group("/api/v1/")
	leads.Get("/", GetLeads)
	leads.Get("/:id", GetLead)
	leads.Post("/", NewLead)
	leads.Delete("/:id", DeleteLead)
	leads.Put("/:id", UpdateLead)
}

func GetLeads(ctx *fiber.Ctx) error {
	db := database.DB
	var leads []types.Lead
	db.Find(&leads).Limit(10)

	return ctx.JSON(leads)
}

func GetLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DB

	var lead types.Lead
	db.First(&lead, "id = ?", id)
	if lead.Name == "" {
		return ctx.Status(500).SendString("No lead found with sent ID")
	}

	return ctx.JSON(lead)
}

func NewLead(ctx *fiber.Ctx) error {
	db := database.DB
	var lead types.Lead

	err := ctx.BodyParser(&lead)
	if err != nil {
		return ctx.Status(503).SendString(err.Error())
	}

	lead.ID = uuid.New()

	db.Create(&lead)

	return ctx.JSON(lead)
}

func DeleteLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DB

	var lead types.Lead
	db.First(&lead, "id = ?", id)
	if lead.Name == "" {
		return ctx.Status(500).SendString("No lead found with sent ID")
	}

	db.Delete(&lead)

	return ctx.SendString("Lead succesfully deleted!")
}

func UpdateLead(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.DB

	var lead types.Lead
	db.First(&lead, "id = ?", id)
	if lead.Name == "" {
		return ctx.Status(500).SendString("No lead found with sent ID")
	}

	var updatedLead types.Lead

	err := ctx.BodyParser(&updatedLead)
	if err != nil {
		return ctx.Status(503).SendString(err.Error())
	}

	db.Model(&lead).Updates(updatedLead)

	return ctx.JSON(lead)
}

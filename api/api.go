package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/murtaza-udaipurwala/sme/db"
)

func setRoutes(app *fiber.App) {
	app.Post("/shorten", handlePOST)
	app.Get("/shorten/:id", handleGET)
}

var database *db.DB

func main() {
	database = db.InitDB()

	app := fiber.New()
	setRoutes(app)
}

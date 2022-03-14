package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/murtaza-udaipurwala/sme/db"
)

func setRoutes(app *fiber.App) {
	app.Post("/shorten", handlePOST)
	app.Get("/:uid", handleGET)
}

var database *db.DB

func Run() {
	database = db.InitDB()

	app := fiber.New()
	setRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

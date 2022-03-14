package api

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/murtaza-udaipurwala/sme/db"
)

func setRoutes(app *fiber.App) {
	app.Post("/shorten", handlePOST)
	app.Get("/:uid", handleGET)
}

var database *db.DB

func Run() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panic(err)
	}

	database = db.InitDB()

	app := fiber.New()
	setRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

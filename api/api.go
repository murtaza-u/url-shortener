package api

import "github.com/gofiber/fiber/v2"

func setRoutes(app *fiber.App) {
	app.Post("/shorten", handlePOST)
	app.Get("/shorten/:id", handleGET)
}

func main() {
	app := fiber.New()
	setRoutes(app)
}

package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/murtaza-udaipurwala/sme/uid"
)

type Request struct {
	URL string `json:"url"`
}

func handlePOST(ctx *fiber.Ctx) error {
	req := new(Request)

	err := ctx.BodyParser(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if !validateURL(req.URL) {
		return ctx.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": errors.New("invalid url"),
		})
	}

	url := enforceHTTP(req.URL)

	uid, err := uid.NewUID(database)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	if err := database.Set(uid, url); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return nil
}

func handleGET(ctx *fiber.Ctx) error {
	return nil
}

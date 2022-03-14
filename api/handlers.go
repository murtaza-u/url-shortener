package api

import (
	"errors"
	"fmt"
	"log"

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
			"error": err.Error(),
		})
	}

	log.Printf("url: %s\n", req.URL)

	if !validateURL(req.URL) {
		return ctx.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"error": errors.New("invalid url").Error(),
		})
	}

	url := enforceHTTP(req.URL)

	uid, err := uid.NewUID(database)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := database.Set(uid, url); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"url": "http://localhost:3000/" + uid,
	})
}

func handleGET(ctx *fiber.Ctx) error {
	uid := ctx.Params("uid")
	log.Printf("uid: %s\n", uid)

	val, err := database.Get(uid)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if val == nil {
		ctx.Status(fiber.StatusNotFound)
		return nil
	}

	return ctx.Redirect(fmt.Sprintf("%s", val), fiber.StatusTemporaryRedirect)
}

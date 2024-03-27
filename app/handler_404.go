package app

import (
	"errors"

	response "github.com/we-and/weand_backend_common/response"

	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		response.SetNotFound(c, "404", errors.New("Not found"))
		return nil
	}
}

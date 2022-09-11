package utils

import (
	"github.com/gofiber/fiber/v2"
	"go-server-jwt/constant"
)

func ResponseMessage(c *fiber.Ctx, message string) error {
	return c.JSON(fiber.Map{
		constant.Message: message,
	})
}

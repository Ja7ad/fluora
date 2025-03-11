package rest

import (
	"github.com/Ja7ad/fluora/internal/errors"
	"github.com/gofiber/fiber/v2"
)

func errorHandlerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			if apiErr, ok := err.(*errors.Errors); ok {
				return c.Status(apiErr.StatusCode).JSON(fiber.Map{
					"code":    apiErr.Code,
					"message": apiErr.Message,
					"details": apiErr.Details,
				})
			}

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": fiber.Map{
					"code":    "INVALID_REQUEST_PATH_ERROR",
					"message": "Invalid request path",
				},
			})
		}

		return nil
	}
}

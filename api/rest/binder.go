package rest

import "github.com/gofiber/fiber/v2"

type Binder interface {
	Register(app *fiber.App, middlewares ...interface{})
}

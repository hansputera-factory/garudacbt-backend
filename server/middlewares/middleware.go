package middlewares

import "github.com/gofiber/fiber/v2"

type UserMiddleware interface {
	LoggedUserAdminOrTeacher(c *fiber.Ctx) error
	LoggedUserAdmin(c *fiber.Ctx) error
	LoggedUserStudent(c *fiber.Ctx) error
	LoggedUserTeacher(c *fiber.Ctx) error
}

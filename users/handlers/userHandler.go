package handlers

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}

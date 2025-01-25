package handlers

import "github.com/gofiber/fiber/v2"

type InstallationHandler interface {
	Install(c *fiber.Ctx) error
	CheckInstall(c *fiber.Ctx) error
}

package handlers

import "github.com/gofiber/fiber/v2"

type SchoolHandler interface {
	CreateSchool(c *fiber.Ctx) error
	ListSchoolShortCodes(c *fiber.Ctx) error
}

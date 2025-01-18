package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"hanifu.id/hansputera-factory/garudacbt-backend/schools/models"
	"hanifu.id/hansputera-factory/garudacbt-backend/schools/usecases"
)

type schoolHttpHandler struct {
	schoolUsecase usecases.SchoolUsecase
}

func NewSchoolHttpHandler(schoolUsecase usecases.SchoolUsecase) SchoolHandler {
	return &schoolHttpHandler{
		schoolUsecase: schoolUsecase,
	}
}

func (h *schoolHttpHandler) ListSchoolShortCodes(c *fiber.Ctx) error {
	return responseWithData(c, fiber.StatusOK, true, h.schoolUsecase.ListSchoolOnlyShortCodes())
}

func (h *schoolHttpHandler) CreateSchool(c *fiber.Ctx) error {
	payload := new(models.AddSchoolModel)

	if err := c.BodyParser(payload); err != nil {
		return response(c, fiber.StatusBadRequest, err.Error(), false)
	}

	if err := h.schoolUsecase.InsertSchool(payload); err != nil {
		return response(c, fiber.StatusInternalServerError, err.Error(), false)
	}

	return response(c, fiber.StatusOK, fmt.Sprintf("School '%s' is created", payload.SchoolName), true)
}

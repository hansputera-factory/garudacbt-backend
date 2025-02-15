package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hanifu.id/hansputera-factory/garudacbt-backend/internal/responses"
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
	return responses.Response(c, fiber.StatusOK, "successfuly fetched", h.schoolUsecase.ListSchoolOnlyShortCodes(), nil)
}

func (h *schoolHttpHandler) CreateSchool(c *fiber.Ctx) error {
	payload := new(models.AddSchoolModel)

	if err := c.BodyParser(payload); err != nil {
		return responses.Response(c, fiber.StatusBadRequest, err.Error(), nil, nil)
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(payload); err != nil {
		return responses.Response(c, fiber.StatusBadRequest, "validation error", nil, err)
	}

	if _, err := h.schoolUsecase.InsertSchool(payload); err != nil {
		return responses.Response(c, fiber.StatusInternalServerError, err.Error(), nil, nil)
	}

	return responses.Response(c, fiber.StatusOK, fmt.Sprintf("School '%s' is created", payload.SchoolName), payload, nil)
}

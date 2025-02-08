package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hanifu.id/hansputera-factory/garudacbt-backend/installations/models"
	"hanifu.id/hansputera-factory/garudacbt-backend/internal/responses"
	schoolUsecases "hanifu.id/hansputera-factory/garudacbt-backend/schools/usecases"
	usersModel "hanifu.id/hansputera-factory/garudacbt-backend/users/models"
	userUsercases "hanifu.id/hansputera-factory/garudacbt-backend/users/usecases"
)

type installationHttpHandler struct {
	schoolUsecase schoolUsecases.SchoolUsecase
	userUsecase   userUsercases.UserUsecase
}

func NewInstallationHttpHandler(schoolUsecase schoolUsecases.SchoolUsecase, userUsecase userUsercases.UserUsecase) InstallationHandler {
	return &installationHttpHandler{
		schoolUsecase: schoolUsecase,
		userUsecase:   userUsecase,
	}
}

func (i *installationHttpHandler) Install(c *fiber.Ctx) error {
	payload := new(models.InsertInstallationModel)
	if err := c.BodyParser(payload); err != nil {
		return responses.Response(c, fiber.StatusBadRequest, err.Error(), nil, nil)
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(payload); err != nil {
		return responses.Response(c, fiber.StatusBadRequest, "validation error", nil, err)
	}

	id, err := i.schoolUsecase.InsertSchool(&payload.School)
	if err != nil {
		return responses.Response(c, fiber.StatusInternalServerError, fmt.Sprintf("School: %s", err.Error()), nil, nil)
	}

	err = i.userUsecase.CreateUser(&usersModel.AddUserModel{
		Name:     payload.User.Username,
		Email:    payload.User.Email,
		Password: payload.User.Password,
		Role:     "admin",
		SchoolID: id,
	})
	if err != nil {
		return responses.Response(c, fiber.StatusInternalServerError, fmt.Sprintf("User: %s", err.Error()), nil, nil)
	}

	return responses.Response(c, fiber.StatusOK, fmt.Sprintf("successfuly install %s", payload.School.SchoolName), payload, nil)
}

func (i *installationHttpHandler) CheckInstall(c *fiber.Ctx) error {
	schools := i.schoolUsecase.ListSchoolOnlyShortCodes()

	if len(schools) > 0 {
		return responses.Response(c, fiber.StatusOK, "successfuly fetched", map[string]int{
			"school_count": len(schools),
		}, nil)
	}

	return responses.Response(c, fiber.StatusNotFound, "no school installed", nil, nil)
}

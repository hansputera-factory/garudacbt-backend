package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/models"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/usecases"
)

type userHttpHandler struct {
	userUsecase usecases.UserUsecase
}

func NewUserHttpHandler(usecase usecases.UserUsecase) UserHandler {
	return &userHttpHandler{
		userUsecase: usecase,
	}
}

func (u *userHttpHandler) LoginUser(c *fiber.Ctx) error {
	payload := new(models.LoginUserModel)

	if err := c.BodyParser(payload); err != nil {
		return response(c, fiber.StatusBadRequest, err.Error(), false)
	}

	payload.ClientIp = c.IP()
	payload.ClientUseragent = c.Get("User-Agent")

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(payload); err != nil {
		return response(c, fiber.StatusBadRequest, err.Error(), false)
	}

	result, err := u.userUsecase.LoginUser(payload)
	if err != nil {
		return response(c, fiber.StatusUnauthorized, err.Error(), false)
	}

	return responseWithData(c, fiber.StatusOK, true, result)
}

func (u *userHttpHandler) CreateUser(c *fiber.Ctx) error {
	payload := new(models.AddUserModel)

	if err := c.BodyParser(payload); err != nil {
		return response(c, fiber.StatusBadRequest, err.Error(), false)
	}

	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(payload); err != nil {
		return response(c, fiber.StatusBadRequest, err.Error(), false)
	}

	if err := u.userUsecase.CreateUser(payload); err != nil {
		return response(c, fiber.StatusInternalServerError, err.Error(), false)
	}

	return response(c, fiber.StatusOK, fmt.Sprintf("user %s successfuly created", payload.Name), true)
}

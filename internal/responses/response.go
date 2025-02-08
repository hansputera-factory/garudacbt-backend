package responses

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Response(c *fiber.Ctx, responseCode int, message string, data any, validationErrors error) error {
	is_ok := true

	if data == nil {
		is_ok = false
	}

	var errors map[string]string = nil
	if validationErrors != nil {
		is_ok = false
		errors = make(map[string]string)
		for _, err := range validationErrors.(validator.ValidationErrors) {
			errors[err.Field()] = err.Error()
		}
	}

	return c.Status(responseCode).JSON(&baseResponse{
		Message:          message,
		Ok:               is_ok,
		Data:             data,
		ValidationErrors: errors,
	})
}

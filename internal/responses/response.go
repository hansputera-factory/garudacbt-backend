package responses

import "github.com/gofiber/fiber/v2"

func Response(c *fiber.Ctx, responseCode int, message string, data any) error {
	is_ok := true

	if data == nil {
		is_ok = false
	}

	return c.Status(responseCode).JSON(&baseResponse{
		Message: message,
		Ok:      is_ok,
		Data:    data,
	})
}

package handlers

import "github.com/gofiber/fiber/v2"

type baseResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Ok         bool   `json:"ok"`
}

func response(c *fiber.Ctx, responseCode int, message string, ok bool) error {
	return c.Status(responseCode).JSON(&baseResponse{
		StatusCode: responseCode,
		Message:    message,
		Ok:         ok,
	})
}

func responseWithData(c *fiber.Ctx, responseCode int, ok bool, data any) error {
	return c.Status(responseCode).JSON(map[string]any{
		"status_code": responseCode,
		"ok":          ok,
		"data":        data,
	})
}

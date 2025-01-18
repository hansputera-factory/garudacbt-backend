package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"hanifu.id/hansputera-factory/garudacbt-backend/internal/claims"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/usecases"
)

type userMiddlewareImpl struct {
	usecase usecases.UserUsecase
}

func NewUserMiddlewareImpl(usecase usecases.UserUsecase) UserMiddleware {
	return &userMiddlewareImpl{
		usecase: usecase,
	}
}

func (u *userMiddlewareImpl) LoggedUserAdminOrTeacher(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*claims.UserClaim)

	if !claims.IsAdmin || !claims.IsTeacher {
		return c.Status(fiber.StatusForbidden).JSON(map[string]any{
			"status_code": fiber.StatusForbidden,
			"message":     "You are not allowed to access this resource",
			"ok":          false,
		})
	}

	return c.Next()
}

func (u *userMiddlewareImpl) LoggedUserAdmin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*claims.UserClaim)

	if !claims.IsAdmin {
		return c.Status(fiber.StatusForbidden).JSON(map[string]any{
			"status_code": fiber.StatusForbidden,
			"message":     "You are not allowed to access this resource",
			"ok":          false,
		})
	}

	return c.Next()
}

func (u *userMiddlewareImpl) LoggedUserTeacher(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*claims.UserClaim)

	if !claims.IsTeacher {
		return c.Status(fiber.StatusForbidden).JSON(map[string]any{
			"status_code": fiber.StatusForbidden,
			"message":     "You are not allowed to access this resource",
			"ok":          false,
		})
	}

	return c.Next()
}

func (u *userMiddlewareImpl) LoggedUserStudent(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*claims.UserClaim)

	if !claims.IsStudent {
		return c.Status(fiber.StatusForbidden).JSON(map[string]any{
			"status_code": fiber.StatusForbidden,
			"message":     "You are not allowed to access this resource",
			"ok":          false,
		})
	}

	return c.Next()
}

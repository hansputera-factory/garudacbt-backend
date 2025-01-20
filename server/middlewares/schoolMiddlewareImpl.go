package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"hanifu.id/hansputera-factory/garudacbt-backend/config"
	"hanifu.id/hansputera-factory/garudacbt-backend/internal/responses"
	"hanifu.id/hansputera-factory/garudacbt-backend/schools/usecases"
)

type schoolMiddlewareImpl struct {
	usecase usecases.SchoolUsecase
	conf    *config.Config
}

func NewSchoolMiddlewareImpl(usecase usecases.SchoolUsecase) SchoolMiddleware {
	return &schoolMiddlewareImpl{
		usecase: usecase,
	}
}

func (s *schoolMiddlewareImpl) OnlyAuthorizedKey(c *fiber.Ctx) error {
	authorizedKeyHeader := c.Get("X-Authorized-Key")

	if strings.Compare(authorizedKeyHeader, s.conf.Secrets.AuthorizeKey) == 0 {
		return c.Next()
	}

	return responses.Response(c, fiber.StatusForbidden, "authorized key isnt pass", nil)
}

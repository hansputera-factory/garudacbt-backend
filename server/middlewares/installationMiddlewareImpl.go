package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"hanifu.id/hansputera-factory/garudacbt-backend/config"
	"hanifu.id/hansputera-factory/garudacbt-backend/internal/responses"
)

type installationMiddlewareImpl struct {
	conf *config.Config
}

func NewInstallationMiddlewareImpl(conf *config.Config) InstallationMiddleware {
	return &installationMiddlewareImpl{
		conf: conf,
	}
}

func (i *installationMiddlewareImpl) OnlyAuthorizedKey(c *fiber.Ctx) error {
	if strings.Compare(c.Get("X-Authorized-Key"), i.conf.Secrets.AuthorizeKey) == 0 {
		return c.Next()
	}

	return responses.Response(c, fiber.StatusUnauthorized, "unauthorized request", nil, nil)
}

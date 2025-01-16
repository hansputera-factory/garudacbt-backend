package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"hanifu.id/hansputera-factory/garudacbt-backend/config"
	"hanifu.id/hansputera-factory/garudacbt-backend/database"
)

type fiberServer struct {
	app  *fiber.App
	db   database.Database
	conf *config.Config
}

func NewFiberServer(conf *config.Config, db database.Database) Server {
	fiberApp := fiber.New(fiber.Config{
		ServerHeader:            "GarudaCBTX",
		JSONEncoder:             sonic.Marshal,
		JSONDecoder:             sonic.Unmarshal,
		EnableTrustedProxyCheck: true,
		EnablePrintRoutes:       true,
		ReduceMemoryUsage:       true,
		AppName:                 "GarudaCBTX-GO",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if strings.Contains(c.Request().URI().String(), "api") {
				return c.Status(code).JSON(map[string]any{
					"message": err.Error(),
					"ok":      false,
				})
			}

			return c.Status(code).SendString(err.Error())
		},
	})

	return &fiberServer{
		app:  fiberApp,
		db:   db,
		conf: conf,
	}
}

func (s *fiberServer) Start() {
	s.app.Use(helmet.New())
	s.app.Use(idempotency.New())
	s.app.Use(cors.New())

	s.app.Get("/fiber_metrics", monitor.New(monitor.Config{
		Title: "GarudaCBTX Metrics",
	}))

	s.initializeSchoolHttpHandler()

	s.app.Listen(fmt.Sprintf("%s:%d", s.conf.Server.Host, s.conf.Server.Port))
}

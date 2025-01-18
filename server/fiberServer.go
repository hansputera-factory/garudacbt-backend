package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/bytedance/sonic"
	jwtware "github.com/gofiber/contrib/jwt"
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

			return c.Status(code).JSON(map[string]any{
				"status_code": code,
				"message":     err.Error(),
				"ok":          false,
			})
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

	s.app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(s.conf.Secrets.JwtKey),
			JWTAlg: "H256",
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusForbidden

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return c.Status(code).JSON(map[string]any{
				"status_code": code,
				"message":     err.Error(),
				"ok":          false,
			})
		},
		Filter: func(c *fiber.Ctx) bool {
			exclusions := []string{
				"GET:/v1/schools",
				"POST:/v1/users/auth",
			}

			for _, exclusion := range exclusions {
				paths := strings.Split(exclusion, ":")
				request_url := strings.TrimSuffix(string(c.Request().URI().Path()), "/")

				if request_url == paths[1] && c.Route().Method == paths[0] {
					return true
				}
			}

			return false
		},
	}))

	// School routes
	s.initializeSchoolHttpHandler()

	// User routes
	s.initializeUserHttpHandler()

	s.app.Listen(fmt.Sprintf("%s:%d", s.conf.Server.Host, s.conf.Server.Port))
}

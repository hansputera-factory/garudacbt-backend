package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/bytedance/sonic"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"hanifu.id/hansputera-factory/garudacbt-backend/config"
	"hanifu.id/hansputera-factory/garudacbt-backend/database"
	"hanifu.id/hansputera-factory/garudacbt-backend/internal/responses"
)

type fiberServer struct {
	app  *fiber.App
	api  fiber.Router
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

			return responses.Response(c, code, err.Error(), nil, nil)
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
	s.app.Use(healthcheck.New(healthcheck.ConfigDefault))

	s.api = s.app.Group("/api")

	s.api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key:    []byte(s.conf.Secrets.JwtKey),
			JWTAlg: "H512",
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusForbidden

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return responses.Response(c, code, err.Error(), nil, nil)
		},
		Filter: func(c *fiber.Ctx) bool {
			exclusions := []string{
				"GET:/api/v1/schools",
				"POST:/api/v1/users/auth",
				"GET:/api/v1/install",
				"POST:/api/v1/install",
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

	s.app.Get("/fiber_metrics", monitor.New(monitor.Config{
		Title: "GarudaCBTX Metrics",
	}))

	// School routes
	s.initializeSchoolHttpHandler()

	// User routes
	s.initializeUserHttpHandler()

	// Installation routes
	s.initializeInstallationHttpHandler()

	s.app.Use("/assets", filesystem.New(filesystem.Config{
		Root: http.Dir("./views/dist/assets"),
	}))
	s.app.Static("*", "./views/dist")

	s.app.Listen(fmt.Sprintf("%s:%d", s.conf.Server.Host, s.conf.Server.Port))
}

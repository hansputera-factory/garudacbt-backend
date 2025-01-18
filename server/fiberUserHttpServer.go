package server

import (
	"hanifu.id/hansputera-factory/garudacbt-backend/server/middlewares"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/handlers"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/repositories"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/usecases"
)

func (f *fiberServer) initializeUserHttpHandler() {
	repository := repositories.NewUserDbRepository(f.db)
	usecase := usecases.NewUserUsecaseImpl(repository)
	httpHandler := handlers.NewUserHttpHandler(usecase)
	userMiddleware := middlewares.NewUserMiddlewareImpl(usecase)

	routers := f.app.Group("/v1/users")

	// WARNING: DON'T FORGET TO CHANGE THE 'mode' IN config.yaml FILE BEFORE MADE IT TO PUBLIC
	// The codes below will allow unauthorized access to do CRUD if 'mode' is set to 'development'
	if f.conf.Mode == "development" {
		routers.Post("/__dev", httpHandler.CreateUser)
	}

	routers.Post("/", userMiddleware.LoggedUserAdmin, httpHandler.CreateUser)

	routers.Post("/auth", httpHandler.LoginUser)
}

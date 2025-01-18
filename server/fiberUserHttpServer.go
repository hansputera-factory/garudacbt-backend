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

	routers.Post("/", userMiddleware.LoggedUserAdmin, httpHandler.CreateUser)

	routers.Post("/auth", httpHandler.LoginUser)
}

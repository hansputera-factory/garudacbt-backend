package server

import (
	schoolHttpHandler "hanifu.id/hansputera-factory/garudacbt-backend/schools/handlers"
	schoolRepositories "hanifu.id/hansputera-factory/garudacbt-backend/schools/repositories"
	schoolUsecase "hanifu.id/hansputera-factory/garudacbt-backend/schools/usecases"
	"hanifu.id/hansputera-factory/garudacbt-backend/server/middlewares"
)

func (f *fiberServer) initializeSchoolHttpHandler() {
	repository := schoolRepositories.NewSchoolDbRepository(f.db)
	usecase := schoolUsecase.NewSchoolUsecaseImpl(repository)
	middleware := middlewares.NewSchoolMiddlewareImpl(usecase)

	httpHandler := schoolHttpHandler.NewSchoolHttpHandler(usecase)

	// Routers
	router := f.api.Group("/v1/schools")
	router.Post("/", middleware.OnlyAuthorizedKey, httpHandler.CreateSchool)
	router.Get("/", httpHandler.ListSchoolShortCodes)
}

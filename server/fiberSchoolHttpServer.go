package server

import (
	schoolHttpHandler "hanifu.id/hansputera-factory/garudacbt-backend/schools/handlers"
	schoolRepositories "hanifu.id/hansputera-factory/garudacbt-backend/schools/repositories"
	schoolUsecase "hanifu.id/hansputera-factory/garudacbt-backend/schools/usecases"
)

func (f *fiberServer) initializeSchoolHttpHandler() {
	repository := schoolRepositories.NewSchoolDbRepository(f.db)
	usecase := schoolUsecase.NewSchoolUsecaseImpl(repository)

	httpHandler := schoolHttpHandler.NewSchoolHttpHandler(usecase)

	// Routers
	router := f.api.Group("/v1/schools")
	router.Post("/", httpHandler.CreateSchool)
	router.Get("/", httpHandler.ListSchoolShortCodes)
}

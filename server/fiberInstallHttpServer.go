package server

import (
	"hanifu.id/hansputera-factory/garudacbt-backend/installations/handlers"
	schoolRepositories "hanifu.id/hansputera-factory/garudacbt-backend/schools/repositories"
	schoolUsecases "hanifu.id/hansputera-factory/garudacbt-backend/schools/usecases"
	"hanifu.id/hansputera-factory/garudacbt-backend/server/middlewares"
	userRepositories "hanifu.id/hansputera-factory/garudacbt-backend/users/repositories"
	userUsecases "hanifu.id/hansputera-factory/garudacbt-backend/users/usecases"
)

func (f *fiberServer) initializeInstallationHttpHandler() {
	userRepository := userRepositories.NewUserDbRepository(f.db)
	userUsecase := userUsecases.NewUserUsecaseImpl(userRepository)

	schoolRepository := schoolRepositories.NewSchoolDbRepository(f.db)
	schoolUsecase := schoolUsecases.NewSchoolUsecaseImpl(schoolRepository)

	httpHandler := handlers.NewInstallationHttpHandler(schoolUsecase, userUsecase)
	middleware := middlewares.NewInstallationMiddlewareImpl(f.conf)

	router := f.api.Group("/v1/install")

	router.Get("/", httpHandler.CheckInstall)
	router.Post("/", middleware.OnlyAuthorizedKey, httpHandler.Install)
}

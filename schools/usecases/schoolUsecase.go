package usecases

import "hanifu.id/hansputera-factory/garudacbt-backend/schools/models"

type SchoolUsecase interface {
	InsertSchool(in *models.AddSchoolModel) error
	ListSchoolOnlyShortCodes() []string
}

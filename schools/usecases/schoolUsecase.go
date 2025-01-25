package usecases

import "hanifu.id/hansputera-factory/garudacbt-backend/schools/models"

type SchoolUsecase interface {
	InsertSchool(in *models.AddSchoolModel) (int64, error)
	ListSchoolOnlyShortCodes() []string
}

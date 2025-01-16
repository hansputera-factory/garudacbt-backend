package usecases

import "hanifu.id/hansputera-factory/garudacbt-backend/schools/models"

type SchoolUsecase interface {
	SchoolDataProcessing(in *models.AddSchoolModel) error
	ListSchoolOnlyShortCodes() []string
}

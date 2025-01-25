package repositories

import "hanifu.id/hansputera-factory/garudacbt-backend/schools/entities"

type SchoolRepository interface {
	InsertSchoolData(in *entities.InsertSchoolDto) (int64, error)
	ListSchoolShortCodes() []string
}

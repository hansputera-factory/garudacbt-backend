package usecases

import (
	"database/sql"

	"hanifu.id/hansputera-factory/garudacbt-backend/schools/entities"
	"hanifu.id/hansputera-factory/garudacbt-backend/schools/models"
	"hanifu.id/hansputera-factory/garudacbt-backend/schools/repositories"
)

type schoolUsecaseImpl struct {
	schoolRepository repositories.SchoolRepository
}

func NewSchoolUsecaseImpl(schoolRepository repositories.SchoolRepository) SchoolUsecase {
	return &schoolUsecaseImpl{
		schoolRepository: schoolRepository,
	}
}

func (s *schoolUsecaseImpl) InsertSchool(in *models.AddSchoolModel) (int64, error) {
	insertPayload := &entities.InsertSchoolDto{
		SchoolName: in.SchoolName,
		ShortCode:  in.ShortCode,
		AppName:    in.AppName,
		Address: sql.NullString{
			String: in.Address,
		},
		Latitude: sql.NullString{
			String: in.Latitude,
		},
		Longitude: sql.NullString{
			String: in.Longitude,
		},
		HeadmasterName: in.HeadmasterName,
		HeadmasterID:   in.HeadmasterID,
		Website: sql.NullString{
			String: in.Website,
		},
		Email: sql.NullString{
			String: in.Email,
		},
	}

	if id, err := s.schoolRepository.InsertSchoolData(insertPayload); err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func (s *schoolUsecaseImpl) ListSchoolOnlyShortCodes() []string {
	return s.schoolRepository.ListSchoolShortCodes()
}

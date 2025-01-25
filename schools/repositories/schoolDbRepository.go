package repositories

import (
	"context"

	"hanifu.id/hansputera-factory/garudacbt-backend/database"
	"hanifu.id/hansputera-factory/garudacbt-backend/schools/entities"
)

type schoolDbRepository struct {
	db  database.Database
	ctx context.Context
}

func NewSchoolDbRepository(db database.Database) SchoolRepository {
	return &schoolDbRepository{db: db, ctx: context.TODO()}
}

func (s *schoolDbRepository) InsertSchoolData(in *entities.InsertSchoolDto) (int64, error) {
	db := s.db.GetDb()

	result, err := db.RegisterSchool(s.ctx, *in)
	if err != nil {
		return 0, err
	}

	if id, err := result.LastInsertId(); err != nil {
		return 0, err
	} else {
		return id, nil
	}
}

func (s *schoolDbRepository) ListSchoolShortCodes() []string {
	db := s.db.GetDb()

	result, err := db.ListSchoolOnlyShortCodes(s.ctx)
	if err != nil {
		return []string{}
	}

	return result
}

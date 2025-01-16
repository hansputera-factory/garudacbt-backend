package entities

import "hanifu.id/hansputera-factory/garudacbt-backend/database"

type (
	InsertSchoolDto = database.RegisterSchoolParams

	InsertSchoolResponse = struct {
		Message string
		Ok      bool
	}
)

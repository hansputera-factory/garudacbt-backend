package entities

import "hanifu.id/hansputera-factory/garudacbt-backend/database"

type (
	InsertUserDto = struct {
		database.RegisterUserParams
		Role string
	}

	GetUserByIdAndAccessDto = struct {
		UserId    int64
		IsStudent bool
		IsAdmin   bool
		IsTeacher bool
	}

	GetUserByNameOrEmailDto = struct {
		User     string
		SchoolID int64
	}

	InsertAuthLogDto = struct {
		UserID          int64
		ClientIp        string
		ClientUseragent string
		SchoolID        int64
	}
)

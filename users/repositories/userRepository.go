package repositories

import (
	"hanifu.id/hansputera-factory/garudacbt-backend/database"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/entities"
)

type UserRepository interface {
	CreateUser(in *entities.InsertUserDto) error
	InsertAuthLog(in *entities.InsertAuthLogDto) error
	GetUserByNameOrEmail(in *entities.GetUserByNameOrEmailDto) (*database.GetUserByNameOrEmailRow, error)
	GetUserByIdAndAccess(in *entities.GetUserByIdAndAccessDto) (*database.GetUserByIdAndAccessRow, error)
}

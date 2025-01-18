package usecases

import (
	"hanifu.id/hansputera-factory/garudacbt-backend/users/models"
)

type UserUsecase interface {
	CreateUser(in *models.AddUserModel) error
	LoginUser(in *models.LoginUserModel) (*models.LoginUserDataModel, error)
	GetUserByIdAndAccess(in *models.CheckUserByIdAndAccessModel) (*models.DataUserModel, error)
}

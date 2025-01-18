package usecases

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
	"hanifu.id/hansputera-factory/garudacbt-backend/config"
	"hanifu.id/hansputera-factory/garudacbt-backend/database"
	"hanifu.id/hansputera-factory/garudacbt-backend/internal/claims"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/entities"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/models"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/repositories"
)

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
	conf           *config.Config
}

func NewUserUsecaseImpl(repository repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{
		userRepository: repository,
	}
}

func (u *userUsecaseImpl) LoginUser(in *models.LoginUserModel) (*models.LoginUserDataModel, error) {
	user, err := u.userRepository.GetUserByNameOrEmail(&entities.GetUserByNameOrEmailDto{
		User:     in.User,
		SchoolID: in.SchoolID,
	})

	if err != nil {
		return nil, err
	}

	if matches, err := argon2.VerifyEncoded([]byte(in.Password), []byte(user.User.Password)); err != nil {
		return nil, err
	} else {
		if !matches {
			return nil, errors.New("incorrect password")
		}
	}

	user_claims := claims.UserClaim{
		UserId:    user.User.ID,
		IsStudent: user.UserAccess.IsStudent == 1,
		IsAdmin:   user.UserAccess.IsAdmin == 1,
		IsTeacher: user.UserAccess.IsTeacher == 1,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "GarudaCBTX",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, user_claims)
	token_encoded, err := token.SignedString([]byte(u.conf.Secrets.JwtKey))
	if err != nil {
		return nil, err
	}

	// create auth log within goroutine
	go u.userRepository.InsertAuthLog(&entities.InsertAuthLogDto{
		UserID:          user.User.ID,
		ClientIp:        in.ClientIp,
		ClientUseragent: in.ClientUseragent,
		SchoolID:        in.SchoolID,
	})

	return &models.LoginUserDataModel{
		Token: token_encoded,
		User: &models.DataUserModel{
			ID:       user.User.ID,
			Name:     user.User.Name,
			Email:    user.User.Email,
			SchoolID: user.User.SchoolID,
			UserAccess: &models.DataUserAccessModel{
				IsStudent: user.UserAccess.IsStudent == 1,
				IsAdmin:   user.UserAccess.IsAdmin == 1,
				IsTeacher: user.UserAccess.IsTeacher == 1,
			},
		},
	}, nil
}

func (u *userUsecaseImpl) GetUserByIdAndAccess(in *models.CheckUserByIdAndAccessModel) (*models.DataUserModel, error) {
	user, err := u.userRepository.GetUserByIdAndAccess(&entities.GetUserByIdAndAccessDto{
		UserId:    in.UserId,
		IsStudent: in.IsStudent,
		IsAdmin:   in.IsAdmin,
		IsTeacher: in.IsTeacher,
	})

	if err != nil {
		return nil, err
	}

	return &models.DataUserModel{
		ID:       user.User.ID,
		Name:     user.User.Name,
		Email:    user.User.Email,
		SchoolID: user.User.SchoolID,
		UserAccess: &models.DataUserAccessModel{
			IsStudent: in.IsStudent,
			IsAdmin:   in.IsAdmin,
			IsTeacher: in.IsTeacher,
		},
	}, nil
}

func (u *userUsecaseImpl) CreateUser(in *models.AddUserModel) error {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(in.Password))
	if err != nil {
		return err
	}

	return u.userRepository.CreateUser(&entities.InsertUserDto{
		RegisterUserParams: database.RegisterUserParams{
			Username: in.Name,
			IsActive: sql.NullInt32{Int32: 0},
			Email:    in.Email,
			SchoolID: in.SchoolID,
			Password: string(encoded),
		},
		Role: in.Role,
	})
}

package repositories

import (
	"context"

	"hanifu.id/hansputera-factory/garudacbt-backend/database"
	"hanifu.id/hansputera-factory/garudacbt-backend/users/entities"
)

type userDbRepository struct {
	db  database.Database
	ctx context.Context
}

func NewUserDbRepository(db database.Database) UserRepository {
	return &userDbRepository{
		db:  db,
		ctx: context.TODO(),
	}
}

func (u *userDbRepository) InsertAuthLog(in *entities.InsertAuthLogDto) error {
	db := u.db.GetDb()

	result, err := db.CreateAuthLog(u.ctx, database.CreateAuthLogParams{
		UserID:          in.UserID,
		ClientIp:        in.ClientIp,
		ClientUseragent: in.ClientUseragent,
		SchoolID:        in.SchoolID,
	})

	if err != nil {
		return nil
	}

	if _, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

func (u *userDbRepository) GetUserByNameOrEmail(in *entities.GetUserByNameOrEmailDto) (*database.GetUserByNameOrEmailRow, error) {
	db := u.db.GetDb()

	result, err := db.GetUserByNameOrEmail(u.ctx, database.GetUserByNameOrEmailParams{
		UserQuery: in.User,
		SchoolID:  in.SchoolID,
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *userDbRepository) GetUserByIdAndAccess(in *entities.GetUserByIdAndAccessDto) (*database.GetUserByIdAndAccessRow, error) {
	db := u.db.GetDb()

	is_teacher := 0
	is_admin := 0
	is_student := 0

	if in.IsStudent {
		is_student = 1
	} else if in.IsAdmin {
		is_admin = 1
	} else if in.IsTeacher {
		is_teacher = 1
	}

	payload := &database.GetUserByIdAndAccessParams{
		UserID:    in.UserId,
		IsStudent: int32(is_student),
		IsAdmin:   int32(is_admin),
		IsTeacher: int32(is_teacher),
	}

	result, err := db.GetUserByIdAndAccess(u.ctx, *payload)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (u *userDbRepository) CreateUser(in *entities.InsertUserDto) error {
	db := u.db.GetDb()
	payload := database.RegisterUserParams{
		IsActive: in.RegisterUserParams.IsActive,
		Username: in.RegisterUserParams.Username,
		Password: in.RegisterUserParams.Password,
		Email:    in.RegisterUserParams.Email,
		SchoolID: in.RegisterUserParams.SchoolID,
	}

	result, err := db.RegisterUser(u.ctx, payload)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// create user access
	is_student := 0
	is_admin := 0
	is_teacher := 0

	if in.Role == "admin" {
		is_admin = 1
	} else if in.Role == "student" {
		is_student = 1
	} else if in.Role == "teacher" {
		is_teacher = 1
	}

	userAccessPayload := &database.CreateUserAccessParams{
		UserID:    userId,
		SchoolID:  in.SchoolID,
		IsStudent: int32(is_student),
		IsAdmin:   int32(is_admin),
		IsTeacher: int32(is_teacher),
	}

	result, err = db.CreateUserAccess(u.ctx, *userAccessPayload)
	if err != nil {
		return err
	}

	if _, err = result.LastInsertId(); err != nil {
		return err
	}

	return nil
}

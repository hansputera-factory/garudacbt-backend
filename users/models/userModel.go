package models

type AddUserModel struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	SchoolID int64  `json:"school_id" validate:"school_id"`
	Role     string `json:"role" validate:"required,oneof=admin student teacher"`
}

type DataUserAccessModel struct {
	IsStudent bool `json:"is_student"`
	IsAdmin   bool `json:"is_admin"`
	IsTeacher bool `json:"is_teacher"`
}

type DataUserModel struct {
	ID         int64                `json:"id"`
	Name       string               `json:"name"`
	Email      string               `json:"email"`
	SchoolID   int64                `json:"school_id"`
	UserAccess *DataUserAccessModel `json:"user_access"`
}

type LoginUserDataModel struct {
	User  *DataUserModel `json:"user"`
	Token string         `json:"token"`
}

type LoginUserModel struct {
	User            string `json:"user" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ClientIp        string `validate:"required,ip"`
	ClientUseragent string `validate:"required"`
	SchoolID        int64  `json:"school_id" validate:"required"`
}

type CheckUserByIdAndAccessModel struct {
	UserId    int64 `json:"user_id" validate:"required"`
	IsStudent bool  `json:"is_student" validate:"required"`
	IsAdmin   bool  `json:"is_admin" validate:"required"`
	IsTeacher bool  `json:"is_teacher" validate:"required"`
}

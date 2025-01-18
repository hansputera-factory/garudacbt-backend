package claims

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	UserId    int64 `json:"user_id"`
	IsStudent bool  `json:"is_student"`
	IsAdmin   bool  `json:"is_admin"`
	IsTeacher bool  `json:"is_teacher"`

	jwt.RegisteredClaims
}

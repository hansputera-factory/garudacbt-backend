package models

type AddSchoolModel struct {
	SchoolName       string `json:"name" validate:"required"`
	ShortCode        string `json:"short_code" validate:"required"`
	SchoolNationalID string `json:"school_national_id" validate:"required"`
	Address          string `json:"address"`
	Latitude         string `json:"latitude" validate:"latitude"`
	Longitude        string `json:"longitude" validate:"longitude"`
	HeadmasterName   string `json:"headmaster_name" validate:"required"`
	HeadmasterID     string `json:"headmaster_id" validate:"required"`
	Website          string `json:"website" validate:"url"`
	Email            string `json:"email" validate:"email"`
	AppName          string `json:"app_name" validate:"required"`
}

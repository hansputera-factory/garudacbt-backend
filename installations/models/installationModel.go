package models

import schoolModels "hanifu.id/hansputera-factory/garudacbt-backend/schools/models"

type InsertInstallationModel struct {
	School schoolModels.AddSchoolModel `json:"school" validate:"required"`
	User   struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	} `json:"user" validate:"required"`
}

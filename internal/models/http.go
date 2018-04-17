package models

import validator "gopkg.in/go-playground/validator.v9"

type (
	TemplateInfo struct {
		Firstname string
		URL       string
	}

	SignupRequest struct {
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required"`
		Firstname string `json:"firstname" validate:"required"`
		Lastname  string `json:"lastname" validate:"required"`
		Birthdate string `json:"birthdate" validate:"required"`
		Gender    string `json:"gender" validate:"required"`
	}

	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginResponse struct {
		*User `json:"user"`
		Token string `json:"token"`
	}

	MatchRequest struct {
		State               string  `json:"state" validate:"required"`
		City                string  `json:"city" validate:"required"`
		Zip                 uint    `json:"zip" validate:"required"`
		HasPets             bool    `json:"hasPets" validate:"required"`
		Smoker              bool    `json:"smoker" validate:"required"`
		TargetOccupantCount uint    `json:"targetOccupantCount" validate:"required"`
		BudgetMax           float32 `json:"budgetMax" validate:"required"`
		BudgetMin           float32 `json:"budgetMin" validate:"required"`
		GenderRequirement   string  `json:"genderRequirement" validate:"required"`
	}

	MatchResponse struct {
		Offer         *RoommateOffer `json:"offer"`
		Compatability float32        `json:"compatability"`
	}

	Validator struct {
		Validator *validator.Validate
	}

	ReportRequest struct {
		Message string `validate:"required"`
	}
)

func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

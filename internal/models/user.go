package models

const (
	AdminLevelUser      = "user"
	AdminLevelModerator = "moderator"
	AdminLevelSusan     = "Susan"
	GenderMale          = "male"
	GenderFemale        = "female"
)

type User struct {
	Email            string `json:"email"`
	Firstname        string `json:"firstname" validate:"required"`
	Lastname         string `json:"lastname" validate:"required"`
	Gender           string `json:"gender" validate:"required"`
	Birthdate        string `json:"birthdate" validate:"required"`
	AdminLevel       string `json:"adminLevel"`
	About            string `json:"about" validate:"required"`
	Tags             []Tag  `json:"tags"`
	ProfileImageUrl  string `json:"profileImageUrl"`
	Active           bool   `json:"active"`
	RegistrationDate string `json:"registrationDate"`
	PasswordHash     string `json:"-"`
}

type Tag struct {
	Id      uint   `json:"id"`
	Content string `json:"content" validate:"required"`
}

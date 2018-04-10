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
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Gender           string `json:"gender"`
	Birthdate        string `json:"birthdate"`
	AdminLevel       string `json:"adminLevel"`
	About            string `json:"about"`
	Tags             []Tag  `json:"tags"`
	ProfileImageUrl  string `json:"profileImageUrl"`
	Active           bool   `json:"active"`
	RegistrationDate string `json:"registrationDate"`
	PasswordHash     string `json:"-"`
}

type Tag struct {
	Id      uint   `json:"id"`
	Content string `json:"content"`
}

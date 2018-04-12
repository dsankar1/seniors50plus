package models

import "github.com/jinzhu/gorm"

const (
	AdminLevelUser      = "user"
	AdminLevelModerator = "moderator"
	AdminLevelSusan     = "susan"
	GenderMale          = "male"
	GenderFemale        = "female"
)

type User struct {
	gorm.Model
	Email            string `json:"email" gorm:"unique; not null"`
	Firstname        string `json:"firstname" validate:"required" gorm:"type:varchar(100); not null"`
	Lastname         string `json:"lastname" validate:"required" gorm:"type:varchar(100); not null"`
	Gender           string `json:"gender" validate:"required" gorm:"type:enum('male','female'); not null"`
	Birthdate        string `json:"birthdate" validate:"required" gorm:"type:date; not null"`
	AdminLevel       string `json:"adminLevel" gorm:"type:enum('user','moderator','susan'); not null; default:'user'"`
	About            string `json:"about" validate:"required" gorm:"type:text; default:\"\""`
	ProfileImageUrl  string `json:"profileImageUrl" gorm:"type:varchar(256)"`
	Active           bool   `json:"active" gorm:"not null; default:false"`
	RegistrationDate string `json:"registrationDate" gorm:"type:timestamp; not null; default:current_timestamp"`
	PasswordHash     string `json:"-" gorm:"type:varchar(256); not null"`
	Tags             []Tag  `json:"tags" gorm:"-"`
}

type Tag struct {
	gorm.Model
	Id      uint   `json:"id"`
	Content string `json:"content" validate:"required"`
}

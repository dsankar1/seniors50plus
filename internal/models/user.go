package models

import (
	"time"
)

const (
	AdminLevelUser      = "user"
	AdminLevelModerator = "moderator"
	AdminLevelSusan     = "susan"
	GenderMale          = "male"
	GenderFemale        = "female"
)

type User struct {
	Email           string    `json:"email" gorm:"primary_key"`
	Firstname       string    `json:"firstname" validate:"required" gorm:"not null"`
	Lastname        string    `json:"lastname" validate:"required" gorm:"not null"`
	Gender          string    `json:"gender" validate:"required" gorm:"type:enum('male','female'); not null"`
	Birthdate       string    `json:"birthdate" validate:"required" gorm:"type:date; not null"`
	AdminLevel      string    `json:"adminLevel" gorm:"type:enum('user','moderator','susan'); default:'user'"`
	About           string    `json:"about" gorm:"type:varchar(3000); default:''"`
	ProfileImageUrl string    `json:"profileImageUrl" gorm:"default:''"`
	Active          bool      `json:"-" gorm:"default:false"`
	PasswordHash    string    `json:"-" gorm:"not null"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Tags            []Tag     `json:"tags" gorm:"foreignkey:Owner"`
}

type Tag struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Owner     string    `json:"-"`
	Content   string    `json:"content" validate:"required"`
}

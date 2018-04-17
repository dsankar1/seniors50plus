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
	ID              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Email           string    `gorm:"unique; not null"`
	Firstname       string    `validate:"required" gorm:"not null"`
	Lastname        string    `validate:"required" gorm:"not null"`
	Gender          string    `validate:"required" gorm:"type:enum('male','female'); not null"`
	Birthdate       string    `validate:"required" gorm:"type:date; not null"`
	AdminLevel      string    `gorm:"type:enum('user','moderator','susan'); not null; default:'user'"`
	About           string    `validate:"required" gorm:"type:varchar(3000); not null"`
	ProfileImageUrl string    `gorm:"not null"`
	Active          bool      `json:"-" gorm:"not null; default:false"`
	Banned          bool      `json:"-" gorm:"not null; default:false"`
	PasswordHash    string    `json:"-" gorm:"not null"`
	Tags            []Tag     `validate:"required" gorm:"foreignkey:UserID"`
	Invitations     []Request `gorm:"foreignkey:UserID"`
	Requests        []Request `gorm:"foreignkey:UserID"`
	Flags           []Flag    `gorm:"foreignkey:UserID"`
}

type Tag struct {
	ID        uint
	CreatedAt time.Time
	UserID    uint   `gorm:"not null"`
	Content   string `validate:"required" gorm:"not null"`
}

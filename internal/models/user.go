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
	Email           string    `json:"email" gorm:"unique; not null"`
	Firstname       string    `json:"firstname" validate:"required" gorm:"not null"`
	Lastname        string    `json:"lastname" validate:"required" gorm:"not null"`
	Gender          string    `json:"gender" validate:"required" gorm:"type:enum('male','female'); not null"`
	Birthdate       string    `json:"birthdate" validate:"required" gorm:"type:date; not null"`
	AdminLevel      string    `json:"adminLevel" gorm:"type:enum('user','moderator','susan'); not null; default:'user'"`
	About           string    `json:"about" validate:"required" gorm:"type:varchar(3000); not null"`
	ProfileImageURL string    `json:"profileImageURL" gorm:"not null"`
	Active          bool      `json:"-" gorm:"not null; default:false"`
	PasswordHash    string    `json:"-" gorm:"not null"`
	Tags            []Tag     `json:"tags" validate:"required" gorm:"foreignkey:UserID"`
	Invitations     []Request `json:"invitations" gorm:"foreignkey:UserID"`
	Requests        []Request `json:"requests" gorm:"foreignkey:UserID"`
	Flags           []Flag    `json:"flags" gorm:"foreignkey:UserID"`
}

type Tag struct {
	ID        uint
	CreatedAt time.Time
	UserID    uint   `json:"userID" gorm:"not null"`
	Content   string `json:"content" validate:"required" gorm:"not null"`
}

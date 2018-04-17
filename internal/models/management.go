package models

import (
	"time"
)

type Report struct {
	ID             uint
	CreatedAt      time.Time
	DeletedAt      *time.Time
	UserID         uint   `json:"userID" gorm:"not null"`
	ReportedUserID uint   `json:"reportedUserID" gorm:"not null"`
	Message        string `json:"message" gorm:"not null"`
}

type Flag struct {
	ID              uint
	CreatedAt       time.Time
	UserID          uint `json:"userID" gorm:"not null"`
	ReportedOfferID uint `json:"reportedOfferID" gorm:"not null"`
}

type Ban struct {
	ID        uint
	CreatedAt time.Time
	DeletedAt *time.Time
	ModID     uint `json:"modID" gorm:"not null"`
	BannedID  uint `json:"bannedID" gorm:"not null"`
}

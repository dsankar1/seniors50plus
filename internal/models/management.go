package models

import (
	"time"
)

type Report struct {
	ID             uint
	CreatedAt      time.Time
	DeletedAt      *time.Time
	UserID         uint   `gorm:"not null"`
	ReportedUserID uint   `gorm:"not null"`
	Message        string `gorm:"not null"`
}

type Flag struct {
	ID              uint
	CreatedAt       time.Time
	UserID          uint `gorm:"not null"`
	ReportedOfferID uint `gorm:"not null"`
}

type Ban struct {
	ID        uint
	CreatedAt time.Time
	DeletedAt *time.Time
	ModID     uint `gorm:"not null"`
	BannedID  uint `gorm:"not null"`
}

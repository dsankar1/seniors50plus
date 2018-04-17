package models

import "time"

const (
	RequestStatusPending  = "pending"
	RequestStatusAccepted = "accepted"
	RequestStatusDenied   = "denied"
)

type Request struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	OfferID   uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Status    string `gorm:"type:enum('pending','accepted','denied'); not null; default:'pending'"`
}

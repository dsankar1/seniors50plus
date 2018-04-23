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
	OfferID   uint          `json:"offerID" gorm:"not null"`
	UserID    uint          `json:"userID" gorm:"not null"`
	Status    string        `json:"status" gorm:"type:enum('pending','accepted','denied'); not null; default:'pending'"`
	Offer     RoommateOffer `json:"offer" gorm:"-"`
	User      User          `json:"user" gorm:"-"`
}

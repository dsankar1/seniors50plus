package models

import (
	"time"
)

const (
	PropertyTypeHouse     = "house"
	PropertyTypeApartment = "apartment"
)

type RoommateOffer struct {
	ID                    uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
	UploaderID            uint    `gorm:"not null"`
	GenderRequirement     string  `validate:"required" gorm:"type:enum('male','female','none'); not null; default:'none'"`
	PreChosenProperty     bool    `gorm:"not null"`
	PropertyType          string  `validate:"required" gorm:"type:enum('house','apartment'); default:'apartment'; not null"`
	State                 string  `validate:"required" gorm:"not null"`
	City                  string  `validate:"required" gorm:"not null"`
	Zip                   uint    `validate:"required" gorm:"not null"`
	Budget                float32 `validate:"required" gorm:"not null"`
	PetsAllowed           bool    `gorm:"not null"`
	SmokingAllowed        bool    `gorm:"not null"`
	TargetResidentCount   uint    `validate:"required" gorm:"not null; default:2"`
	AcceptedResidentCount uint    `gorm:"not null; default:1"`
	PropertyImageUrl      string
	Residents             []Resident             `gorm:"foreignkey:OfferID"`
	Requests              []CommunicationRequest `gorm:"foreignkey:OfferID"`
}

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
	UploaderID            uint      `json:"uploaderID" gorm:"not null"`
	GenderRequirement     string    `json:"genderRequirement" validate:"required" gorm:"type:enum('male','female','none'); not null; default:'none'"`
	PreChosenProperty     bool      `json:"preChosenProperty" gorm:"not null"`
	PropertyType          string    `json:"propertyType" validate:"required" gorm:"type:enum('house','apartment'); default:'apartment'; not null"`
	State                 string    `json:"state" validate:"required" gorm:"not null"`
	City                  string    `json:"city" validate:"required" gorm:"not null"`
	Zip                   uint      `json:"zip" validate:"required" gorm:"not null"`
	Budget                float32   `json:"budget" validate:"required" gorm:"not null"`
	PetsAllowed           bool      `json:"petsAllowed" gorm:"not null"`
	Bathrooms             uint      `json:"bathrooms" gorm:"not null; default: 2"`
	Bedrooms              uint      `json:"bedrooms" gorm:"not null; default: 2"`
	SmokingAllowed        bool      `json:"smokingAllowed" gorm:"not null"`
	TargetResidentCount   uint      `json:"targetResidentCount" validate:"required" gorm:"not null; default:2"`
	AcceptedResidentCount uint      `json:"acceptedResidentCount" gorm:"not null; default:1"`
	PropertyImageURL      string    `json:"propertyImageURL"`
	Residents             []Request `json:"residents" gorm:"foreignkey:OfferID"`
	Requests              []Request `json:"requests" gorm:"foreignkey:OfferID"`
}

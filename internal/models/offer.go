package models

import "github.com/jinzhu/gorm"

const (
	PropertyTypeHouse     = "house"
	PropertyTypeApartment = "apartment"
)

type RoommateOffer struct {
	gorm.Model
	PostedBy            User       `json:"postedBy" gorm:"foreignkey:PosterEmail"`
	PosterEmail         string     `json:"-" gorm:"not null; unique"`
	GenderRequirement   string     `json:"genderRequirement" validate:"required" gorm:"type:enum('male','female','none'); not null; default:'none'"`
	PreChosenProperty   bool       `json:"preChosenProperty" validate:"required" gorm:"not null"`
	State               string     `json:"state" validate:"required" gorm:"not null"`
	City                string     `json:"city" validate:"required" gorm:"not null"`
	Zip                 uint       `json:"zip" validate:"required" gorm:"not null"`
	Budget              float32    `json:"budget" validate:"required" gorm:"not null"`
	PetsAllowed         bool       `json:"petsAllowed" validate:"required" gorm:"not null"`
	SmokingAllowed      bool       `json:"smokingAllowed" validate:"required" gorm:"not null"`
	TargetOccupantCount uint       `json:"targetOccupantCount" validate:"required" gorm:"not null; default:2"`
	PropertyImageUrl    string     `json:"propertyImageUrl"`
	PropertyType        string     `json:"propertyType" validate:"required" gorm:"not null"`
	Residents           []Resident `json:"residents" gorm:"foreignkey:OfferId"`
	//Requests            []OthersRequest `json:"requests"`
}

type Resident struct {
	gorm.Model
	Occupant      User   `json:"occupant" gorm:"foreignkey:OccupantEmail"`
	OccupantEmail string `json:"-" gorm:"not null"`
	OfferId       uint   `json:"offerId" gorm:"not null"`
}

/*var ExampleOffer = RoommateOffer{
	GenderRequirement:   GenderMale,
	PreChosenProperty:   false,
	State:               "Georgia",
	City:                "Marietta",
	Zip:                 30008,
	BudgetMax:           1100.00,
	BudgetMin:           900.00,
	PetsAllowed:         true,
	SmokingAllowed:      false,
	TargetOccupantCount: 2,
	PropertyImageUrl:    "https://someamazonbucket.com",
}*/

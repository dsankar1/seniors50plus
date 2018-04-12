package models

const (
	PropertyTypeHouse     = "house"
	PropertyTypeApartment = "apartment"
)

type RoommateOffer struct {
	Id                  uint            `json:"id"`
	PostedBy            string          `json:"postedBy"`
	GenderRequirement   string          `json:"genderRequirement" validate:"required"`
	PreChosenProperty   bool            `json:"preChosenProperty" validate:"required"`
	State               string          `json:"state" validate:"required"`
	City                string          `json:"city" validate:"required"`
	Zip                 uint            `json:"zip" validate:"required"`
	Budget              float32         `json:"budget" validate:"required"`
	PetsAllowed         bool            `json:"petsAllowed" validate:"required"`
	SmokingAllowed      bool            `json:"smokingAllowed" validate:"required"`
	TargetOccupantCount uint            `json:"targetOccupantCount" validate:"required"`
	PropertyImageUrl    string          `json:"propertyImageUrl"`
	PostedOn            string          `json:"postedOn"`
	PropertyType        string          `json:"propertyType" validate:"required"`
	Occupants           []Occupant      `json:"occupants"`
	Requests            []OthersRequest `json:"requests"`
}

type Occupant struct {
	Id              uint   `json:"id"`
	Email           string `json:"email"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Gender          string `json:"gender"`
	ProfileImageUrl string `json:"profileImageUrl"`
	AcceptedOn      string `json:"acceptedOn"`
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

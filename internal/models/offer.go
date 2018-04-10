package models

type RoommateOffer struct {
	Id                  uint64   `json:"id"`
	PostedBy            string   `json:"postedBy"`
	GenderRequirement   string   `json:"genderRequirement"`
	PreChosenProperty   bool     `json:"preChosenProperty"`
	State               string   `json:"state"`
	City                string   `json:"city"`
	Zip                 uint     `json:"zip"`
	BudgetMax           float32  `json:"budgetMax"`
	BudgetMin           float32  `json:"budgetMin"`
	PetsAllowed         bool     `json:"petsAllowed"`
	SmokingAllowed      bool     `json:"smokingAllowed"`
	Occupants           []string `json:"occupants"`
	TargetOccupantCount uint     `json:"targetOccupantCount"`
	PropertyImageUrl    string   `json:"propertyImageUrl"`
	Active              bool     `json:"active"`
}

var ExampleOffer = RoommateOffer{
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
	Active:              true,
}

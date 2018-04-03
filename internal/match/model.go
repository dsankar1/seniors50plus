package match

import "seniors50plus/internal/user"

type MatchRequest struct {
	Location            `json:"location"`
	HasPets             bool    `json:"hasPets"`
	Smoker              bool    `json:"smoker"`
	TargetOccupantCount uint    `json:"targetOccupantCount"`
	BudgetMax           float32 `json:"budgetMax"`
	BudgetMin           float32 `json:"budgetMin"`
}

type Match struct {
	Offer         RoommateOffer `json:"offer"`
	Compatability float32       `json:"compatability"`
}

type RoommateOffer struct {
	Id                  uint64 `json:"id"`
	PostedBy            string `json:"postedBy"`
	GenderRequirement   string `json:"genderRequirement"`
	PreChosenProperty   bool   `json:"preChosenProperty"`
	Location            `json:"location"`
	BudgetMax           float32  `json:"budgetMax"`
	BudgetMin           float32  `json:"budgetMin"`
	PetsAllowed         bool     `json:"petsAllowed"`
	SmokingAllowed      bool     `json:"smokingAllowed"`
	Occupants           []string `json:"occupants"`
	TargetOccupantCount uint     `json:"targetOccupantCount"`
	PropertyImageUrl    string   `json:"propertyImageUrl"`
	Active              bool     `json:"active"`
}

type Location struct {
	State string `json:"state"`
	City  string `json:"city"`
	Zip   uint   `json:"zip"`
}

var ExampleOffer = RoommateOffer{
	GenderRequirement: user.Male,
	PreChosenProperty: false,
	Location: Location{
		State: "Georgia",
		City:  "Marietta",
		Zip:   30008,
	},
	BudgetMax:           1100.00,
	BudgetMin:           900.00,
	PetsAllowed:         true,
	SmokingAllowed:      false,
	TargetOccupantCount: 2,
	PropertyImageUrl:    "https://someamazonbucket.com",
	Active:              true,
}

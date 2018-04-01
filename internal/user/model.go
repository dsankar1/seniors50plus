package user

type User struct {
	Email            string   `json:"email"`
	Firstname        string   `json:"firstname"`
	Lastname         string   `json:"lastname"`
	Gender           string   `json:"gender"`
	Age              uint16   `json:"age"`
	Budget           float32  `json:"budget"`
	Smoker           bool     `json:"smoker"`
	PetOwner         bool     `json:"petOwner"`
	MaxRoommateCount uint     `json:"maxRoommateCount"`
	SeekingWithCount uint     `json:"seekingWithCount"`
	LocationOwner    bool     `json:"locationOwner"`
	Active           bool     `json:"active"`
	Admin            bool     `json:"admin"`
	Blocked          []string `json:"blocked"`
	Tags             []string `json:"tags"`
	Location         `json:"location"`
}

type Location struct {
	State string `json:"state"`
	City  string `json:"city"`
	Zip   uint   `json:"zip"`
}

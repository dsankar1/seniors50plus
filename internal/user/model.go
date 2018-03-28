package user

type User struct {
	Email     string  `json:"email"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Gender    string  `json:"gender"`
	Age       uint16  `json:"age"`
	Budget    float32 `json:"budget"`
	Location  `json:"location"`
	Tags      []string `json:"tags"`
	Admin     bool     `json:"admin"`
}

type Location struct {
	State string `json:"state"`
	City  string `json:"city"`
	Zip   uint   `json:"zip"`
}

package user

const (
	AdminLevelUser      = "user"
	AdminLevelModerator = "moderator"
	AdminLevelSusan     = "Susan"
	Male                = "male"
	Female              = "female"
	Other               = "other"
)

type User struct {
	Email           string   `json:"email"`
	Firstname       string   `json:"firstname"`
	Lastname        string   `json:"lastname"`
	Gender          string   `json:"gender"`
	Age             uint16   `json:"age"`
	AdminLevel      string   `json:"adminLevel"`
	About           string   `json:"about"`
	Tags            []string `json:"tags"`
	ProfileImageUrl string   `json:"profileImageUrl"`
}

var ExampleUser = User{
	Email:           "JohnDoe@gmail.com",
	Firstname:       "John",
	Lastname:        "Doe",
	Gender:          "male",
	Age:             56,
	AdminLevel:      AdminLevelUser,
	About:           "Hello, I'm John Doe!",
	Tags:            []string{"Neat freak", "Early Bird", "Vegan", "Sports", "Dogs"},
	ProfileImageUrl: "https://someamazonbucket.com",
}

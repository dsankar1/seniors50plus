# Roommates 40 Plus Server/API

Roommates 40 plus is a roommate finder targeted for users of age 40 or older. This is the web server and Rest API of the app. Below will give an overview of what the API has to offer, including the different endpoints, methods and request/response models.

## Models
```
type User struct {
	ID              uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Email           string   
	Firstname       string    
	Lastname        string    
	Gender          string    
	Birthdate       string    
	AdminLevel      string    
	About           string    
	ProfileImageURL string    
	Active          bool      
	PasswordHash    string    
	Tags            []Tag     
	Invitations     []Request 
	Requests        []Request 
	Flags           []Flag    
}

type Tag struct {
	ID        uint
	CreatedAt time.Time
	UserID    uint   
	Content   string 
}

type RoommateOffer struct {
	ID                    uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             *time.Time
	UploaderID            uint      
	GenderRequirement     string    
	PreChosenProperty     bool      
	PropertyType          string    
	State                 string    
	City                  string   
	Zip                   uint      
	Budget                float32   
	PetsAllowed           bool     
	SmokingAllowed        bool      
	TargetResidentCount   uint      
	AcceptedResidentCount uint      
	PropertyImageURL      string    
	Residents             []Request 
	Requests              []Request 
}

type Request struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	OfferID   uint   
	UserID    uint   
	Status    string 
}

type Report struct {
	ID             uint
	CreatedAt      time.Time
	DeletedAt      *time.Time
	UserID         uint   `json:"userID" gorm:"not null"`
	ReportedUserID uint   `json:"reportedUserID" gorm:"not null"`
	Message        string `json:"message" gorm:"not null"`
}

type Flag struct {
	ID              uint
	CreatedAt       time.Time
	UserID          uint 
	ReportedOfferID uint 
}

type Ban struct {
	ID        uint
	CreatedAt time.Time
	DeletedAt *time.Time
	ModID     uint 
	BannedID  uint 
}
```

## Endpoints

Description: Returns all user information along with an resource token
Endpoint: /api/auth/login
Method: POST
Request Body:
```
{
    "email":"",
    "password":""
}
```

Description: Registers a new inactive user account (confirm email to activate)
Endpoint: /api/auth/signup
Method: POST
Request Body:
```
{
    "email":"",
    "password":"",
    "firstname":"",
    "lastname":"",
    "gender":"male" or "female",
    "birthdate":"YYYY-MM-DD"
}
```

### *MUST INCLUDE RESOURCE TOKEN FOR FOLLOWING ENDPOINTS*

Description: Get the details of a user with :id (Excludes email)
Endpoint: /api/user/:id
Method: GET
Request Body: none

Description: Get the email of the user with :id
*note - only available if communication request is accepted*	
Endpoint: /api/user/:id/email
Method: GET
Request Body: none

Description: Get a detailed list of users provided a list of emails
Endpoint: /api/user/list
Method: POST
Request Body: 
```
[
    {
        "id":""
    },
    ...
]
```

// Gets information on the currently logged in user
Endpoint: /api/user
Method: GET
Request Body: none

// Edits the information of the currently logged in user
*note - Must include all fields, even if they aren't changed
Endpoint: /api/user
Method: POST
Request Body: 
{
    "firstname":"",
    "lastname":"",
    "gender":"male" or "female",
    "birthdate":"YYYY-MM-DD",
    "about":"",
    "tags": [
        {
            "content":""
        }, 
        ...
    ]
}

## Built With

* [Echo](https://echo.labstack.com/guide) - The web framework used
* [gorm](https://github.com/jinzhu/gorm) - ORM used
* [godep](https://github.com/tools/godep) - Dependency management tool

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details



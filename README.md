# Roommates 40 Plus Server/API

Roommates 40 plus is a roommate finder targeted for users of age 40 or older. This is the web server and Rest API of the app. Below will give an overview of what the API has to offer, including the different endpoints, methods and request/response models.

## Models
```go
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

```go
// AUTH ENDPOINTS
e.POST("/api/auth/login", auth.LoginHandler)

e.POST("/api/auth/signup", auth.SignupHandler)

e.GET("/api/auth/signup/confirmation", auth.EmailConfirmationHandler)

// USER ENDPOINTS
e.GET("/api/user/:id", user.GetUserHandler)

e.GET("/api/user/:id/email", user.GetUserEmailHandler)

e.POST("/api/user/list", user.GetUserListHandler)

e.GET("/api/user", user.GetMyselfHandler)

e.PUT("/api/user", user.UpdateUserHandler)

// OFFERS ENDPOINTS
e.POST("/api/offer", offer.PostOfferHandler)

e.GET("/api/offer", offer.GetMyOfferHandler)

e.GET("/api/offer/:id", offer.GetOfferHandler)

e.GET("/api/offer/:id/email", offer.GetOfferEmailHandler)

e.DELETE("/api/offer", offer.DeleteMyOfferHandler)

// COMMUNICATION ENDPOINTS
e.POST("/api/offer/:id/request", request.CreateCommunicationRequestHandler)

e.DELETE("/api/offer/:id/request", request.DeleteCommunicationRequestHandler)

e.PUT("/api/offer/request/:id", request.RespondToCommunicationRequestHandler) //?status=value

// RESIDENT ENDPOINTS
e.POST("/api/user/:id/request", request.CreateResidentRequestHandler)

e.DELETE("/api/user/:id/request", request.DeleteResidentRequestHandler)

e.PUT("/api/user/request/:id", request.RespondToResidentRequestHandler) //?status=value

// REPORT USER ENDPOINTS
e.POST("/api/user/:id/report", management.ReportUserHandler)

e.DELETE("/api/user/:id/report", management.ResolveReportsHandler)

// FLAG OFFER ENDPOINTS
e.POST("/api/offer/:id/flag", management.FlagOfferHandler)

e.DELETE("/api/offer/:id/flag", management.UnflagOfferHandler)

// MOD ENDPOINTS
e.GET("/api/user/report", management.GetReportsHandler)

e.GET("/api/offer/flag", management.GetFlaggedOffers)

e.POST("/api/user/:id/ban", management.BanUserHandler)

e.DELETE("/api/user/:id/ban", management.UnbanUserHandler)

e.GET("/api/user/ban", management.GetBannedUsersHandler)

e.POST("/api/offer/:id/ban", management.BanOfferHandler)

// SUSAN ENDPOINTS
e.GET("/api/user/mod", management.GetModsHandler)

e.POST("/api/user/:id/mod", management.ModUserHandler)

e.DELETE("/api/user/:id/mod", management.UnmodUserHandler)
```

## Built With

* [Echo](https://echo.labstack.com/guide) - The web framework used
* [gorm](https://github.com/jinzhu/gorm) - ORM used
* [godep](https://github.com/tools/godep) - Dependency management tool

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details



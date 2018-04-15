package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnection struct {
	dbstring string
}

func NewDatabaseConnection() *DatabaseConnection {
	dbstring := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v", "capstoneuser",
		os.Getenv("DB_PASSWORD"),
		"capstone.cczajq2nppkf.us-east-2.rds.amazonaws.com",
		"roommates40plusv2?charset=utf8&parseTime=true",
	)
	return &DatabaseConnection{dbstring}
}

func (dbc *DatabaseConnection) CreateTable(models ...interface{}) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.CreateTable(models...).Error
	}
}

// User related DB methods
func (dbc *DatabaseConnection) CreateUser(user *User) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Create(user).Error
	}
}

func (dbc *DatabaseConnection) GetUser(user *User) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Where(user).First(user).Error
	}
}

func (dbc *DatabaseConnection) UpdateUser(user *User) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		if err := db.Save(user).Error; err != nil {
			return err
		}
		if err := db.Where("user_id = ?", user.ID).Delete(&Tag{}).Error; err != nil {
			return err
		}
		for i := 0; i < len(user.Tags); i++ {
			user.Tags[i].UserID = user.ID
			if err := db.Save(&user.Tags[i]).Error; err != nil {
				return err
			}
		}
		return nil
	}
}

func (dbc *DatabaseConnection) AttachTags(user *User) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Where("user_id = ?", user.ID).Find(&user.Tags).Error
	}
}

// Offer related DB methods
func (dbc *DatabaseConnection) CreateOffer(offer *RoommateOffer) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		if results := db.Where(offer).First(&RoommateOffer{}); results.Error != nil {
			if results.RecordNotFound() {
				return db.Create(offer).Error
			}
			return results.Error
		}
		return errors.New("Offer already exists")
	}
}

func (dbc *DatabaseConnection) DeleteOffer(offer *RoommateOffer) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Delete(offer).Error
	}
}

func (dbc *DatabaseConnection) GetOffer(offer *RoommateOffer) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Where(offer).First(offer).Error
	}
}

func (dbc *DatabaseConnection) AttachResidents(offer *RoommateOffer) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Where("offer_id = ?", offer.ID).Find(&offer.Residents).Error
	}
}

func (dbc *DatabaseConnection) AttachRequests(offer *RoommateOffer) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Where("offer_id = ?", offer.ID).Find(&offer.Requests).Error
	}
}

// Communication request related db methods
func (dbc *DatabaseConnection) GetCommunicationRequest(request *Request) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Table("communication_requests").Where(request).First(request).Error
	}
}

func (dbc *DatabaseConnection) CreateCommunicationRequest(request *Request) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		offer := RoommateOffer{
			ID: request.OfferID,
		}
		if err := dbc.GetOffer(&offer); err != nil {
			return err
		}
		if offer.UploaderID == request.UserID {
			return errors.New("You can't request your own post")
		}
		if results := db.Table("communication_requests").Where(request).First(&Request{}); results.Error != nil {
			if results.RecordNotFound() {
				return db.Table("communication_requests").Create(request).Error
			}
			return results.Error
		}
		return errors.New("Request already exists")
	}
}

func (dbc *DatabaseConnection) UpdateCommunicationRequest(request *Request) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		return db.Table("communication_requests").Save(request).Error
	}
}

func (dbc *DatabaseConnection) DeleteCommunicationRequest(request *Request) error {
	if db, err := gorm.Open("mysql", dbc.dbstring); err != nil {
		return err
	} else {
		defer db.Close()
		if err := dbc.GetCommunicationRequest(request); err != nil {
			return err
		}
		if request.Status == RequestStatusDenied {
			return errors.New("Can't delete denied requests")
		}
		return db.Table("communication_requests").Delete(request).Error
	}
}

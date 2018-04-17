package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnection struct {
	User         string
	Password     string
	Endpoint     string
	DatabaseName string
	db           *gorm.DB
}

func NewDatabaseConnection() *DatabaseConnection {
	return &DatabaseConnection{
		User:         "capstoneuser",
		Password:     os.Getenv("DB_PASSWORD"),
		Endpoint:     "capstone.cczajq2nppkf.us-east-2.rds.amazonaws.com",
		DatabaseName: "roommates40plusv2",
	}
}

func (dbc *DatabaseConnection) Open() error {
	dbstring := fmt.Sprintf(
		"%v:%v@tcp(%v)/%v?charset=utf8&parseTime=true",
		dbc.User,
		dbc.Password,
		dbc.Endpoint,
		dbc.DatabaseName,
	)
	var err error
	if dbc.db, err = gorm.Open("mysql", dbstring); err != nil {
		return err
	}
	return nil
}

func (dbc *DatabaseConnection) Close() error {
	return dbc.db.Close()
}

func (dbc *DatabaseConnection) CreateTable(models ...interface{}) error {
	return dbc.db.CreateTable(models...).Error
}

// User related DB methods
func (dbc *DatabaseConnection) CreateUser(user *User) error {
	db := dbc.db
	return db.Create(user).Error
}

func (dbc *DatabaseConnection) GetUser(user *User) error {
	db := dbc.db
	return db.Where(user).First(user).Error
}

func (dbc *DatabaseConnection) GetMods(users *[]User) error {
	db := dbc.db
	user := User{
		AdminLevel: AdminLevelModerator,
	}
	return db.Where(user).Find(users).Error
}

func (dbc *DatabaseConnection) UpdateUser(user *User) error {
	db := dbc.db
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

func (dbc *DatabaseConnection) AttachTags(user *User) error {
	db := dbc.db
	return db.Where("user_id = ?", user.ID).Find(&user.Tags).Error
}

func (dbc *DatabaseConnection) AttachCommunicationRequests(user *User) error {
	db := dbc.db
	return db.Table("communication_requests").Where("user_id = ?", user.ID).Find(&user.Requests).Error
}

func (dbc *DatabaseConnection) AttachResidentInvitations(user *User) error {
	db := dbc.db
	return db.Table("residents").Where("user_id = ?", user.ID).Find(&user.Invitations).Error
}

func (dbc *DatabaseConnection) AttachFlags(user *User) error {
	db := dbc.db
	return db.Where("user_id = ?", user.ID).Find(&user.Flags).Error
}

// Offer related DB methods
func (dbc *DatabaseConnection) CreateOffer(offer *RoommateOffer) error {
	db := dbc.db
	if results := db.Where(offer).First(&RoommateOffer{}); results.Error != nil {
		if results.RecordNotFound() {
			return db.Create(offer).Error
		}
		return results.Error
	}
	return errors.New("Offer already exists")
}

func (dbc *DatabaseConnection) RemovePendingResidentRequests(offer *RoommateOffer) error {
	db := dbc.db
	return db.Table("residents").Where("offer_id like ?", offer.ID).Where("status like ?", RequestStatusPending).Delete(Request{}).Error
}

func (dbc *DatabaseConnection) UpdateOffer(offer *RoommateOffer) error {
	db := dbc.db
	return db.Save(offer).Error
}

func (dbc *DatabaseConnection) DeleteOffer(offer *RoommateOffer) error {
	db := dbc.db
	if err := db.Where(offer).First(offer).Error; err != nil {
		return err
	}
	if err := db.Table("communication_requests").Where("offer_id like ?", offer.ID).Delete(Request{}).Error; err != nil {
		return err
	}
	if err := db.Table("residents").Where("offer_id like ?", offer.ID).Delete(Request{}).Error; err != nil {
		return err
	}
	if err := db.Where("reported_offer_id like ?", offer.ID).Delete(Flag{}).Error; err != nil {
		return err
	}
	return db.Delete(offer).Error
}

func (dbc *DatabaseConnection) GetOffer(offer *RoommateOffer) error {
	db := dbc.db
	return db.Where(offer).First(offer).Error
}

func (dbc *DatabaseConnection) AttachResidents(offer *RoommateOffer) error {
	db := dbc.db
	return db.Table("residents").Where("offer_id = ?", offer.ID).Find(&offer.Residents).Error
}

func (dbc *DatabaseConnection) AttachRequests(offer *RoommateOffer) error {
	db := dbc.db
	return db.Table("communication_requests").Where("offer_id = ?", offer.ID).Find(&offer.Requests).Error
}

// Communication request related db methods
func (dbc *DatabaseConnection) GetCommunicationRequest(request *Request) error {
	db := dbc.db
	return db.Table("communication_requests").Where(request).First(request).Error
}

func (dbc *DatabaseConnection) CreateCommunicationRequest(request *Request) error {
	db := dbc.db
	if results := db.Table("communication_requests").Where(request).First(&Request{}); results.Error != nil {
		if results.RecordNotFound() {
			return db.Table("communication_requests").Create(request).Error
		}
		return results.Error
	}
	return errors.New("Request already exists")
}

func (dbc *DatabaseConnection) UpdateCommunicationRequest(request *Request) error {
	db := dbc.db
	return db.Table("communication_requests").Save(request).Error
}

func (dbc *DatabaseConnection) DeleteCommunicationRequest(request *Request) error {
	db := dbc.db
	/*if err := dbc.GetCommunicationRequest(request); err != nil {
		return err
	}
	if request.Status == RequestStatusDenied {
		return errors.New("Can't delete denied requests")
	}*/
	return db.Table("communication_requests").Delete(request).Error
}

// Resident request related db methods
func (dbc *DatabaseConnection) GetResidentRequest(request *Request) error {
	db := dbc.db
	return db.Table("residents").Where(request).First(request).Error
}

func (dbc *DatabaseConnection) CreateResidentRequest(request *Request) error {
	db := dbc.db
	if results := db.Table("residents").Where(request).First(&Request{}); results.Error != nil {
		if results.RecordNotFound() {
			return db.Table("residents").Create(request).Error
		}
		return results.Error
	}
	return errors.New("Request already exists")
}

func (dbc *DatabaseConnection) UpdateResidentRequest(request *Request) error {
	db := dbc.db
	return db.Table("residents").Save(request).Error
}

func (dbc *DatabaseConnection) DeleteResidentRequest(request *Request) error {
	db := dbc.db
	/*if err := dbc.GetResidentRequest(request); err != nil {
		return err
	}
	if request.Status == RequestStatusDenied {
		return errors.New("Can't delete denied requests")
	}*/
	return db.Table("residents").Delete(request).Error
}

// Flag related db methods
func (dbc *DatabaseConnection) CreateFlag(flag *Flag) error {
	db := dbc.db
	if results := db.Where(flag).First(&Flag{}); results.Error != nil {
		if results.RecordNotFound() {
			return db.Create(flag).Error
		}
		return results.Error
	}
	return errors.New("Already flagged")
}

func (dbc *DatabaseConnection) DeleteFlag(flag *Flag) error {
	db := dbc.db
	return db.Delete(flag).Error
}

func (dbc *DatabaseConnection) GetAllFlags(flags *[]Flag) error {
	db := dbc.db
	return db.Find(flags).Error
}

// Report related db methods
func (dbc *DatabaseConnection) CreateReport(report *Report) error {
	db := dbc.db
	return db.Create(report).Error
}

func (dbc *DatabaseConnection) DeleteReports(userId uint) error {
	db := dbc.db
	return db.Where("reported_user_id like ?", userId).Delete(Report{}).Error
}

func (dbc *DatabaseConnection) GetAllReports(reports *[]Report) error {
	db := dbc.db
	return db.Find(reports).Error
}

// Ban related db methods
func (dbc *DatabaseConnection) GetAllBans(bans *[]Ban) error {
	db := dbc.db
	return db.Find(bans).Error
}

func (dbc *DatabaseConnection) CreateBan(ban *Ban) error {
	db := dbc.db
	user := User{
		ID: ban.BannedID,
	}
	if err := dbc.GetUser(&user); err != nil {
		return err
	}
	if results := db.Where(ban).First(&Ban{}); results.Error != nil {
		if results.RecordNotFound() {
			if err := db.Where("user_id like ?", ban.BannedID).Delete(Tag{}).Error; err != nil {
				return err
			}
			if err := db.Table("communication_requests").Where("user_id like ?", ban.BannedID).Delete(Request{}).Error; err != nil {
				return err
			}
			if err := db.Table("residents").Where("user_id like ?", ban.BannedID).Delete(Request{}).Error; err != nil {
				return err
			}
			if err := db.Where("user_id like ?", ban.BannedID).Delete(Flag{}).Error; err != nil {
				return err
			}
			offer := RoommateOffer{
				UploaderID: ban.BannedID,
			}
			if err := dbc.DeleteOffer(&offer); err != nil && err.Error() != "record not found" {
				return err
			}
			user.Banned = true
			if err := db.Save(&user).Error; err != nil {
				return err
			}
			return db.Create(ban).Error
		}
		return results.Error
	}
	return errors.New("Already banned")
}

func (dbc *DatabaseConnection) DeleteBan(ban *Ban) error {
	db := dbc.db
	user := User{
		ID: ban.BannedID,
	}
	if err := dbc.GetUser(&user); err != nil {
		return err
	}
	user.Banned = false
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return db.Delete(ban).Error
}

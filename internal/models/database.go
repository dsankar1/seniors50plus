package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConnection struct {
	User         string
	Password     string
	Endpoint     string
	DatabaseName string
}

func NewDatabaseConnection() *DatabaseConnection {
	user := "capstoneuser"
	password := "capst0ne18project!" //os.Getenv("DB_PASSWORD")
	endpoint := "capstone.cczajq2nppkf.us-east-2.rds.amazonaws.com"
	databaseName := "roommates40plus"
	return &DatabaseConnection{
		User:         user,
		Password:     password,
		Endpoint:     endpoint,
		DatabaseName: databaseName,
	}
}

func (dbc *DatabaseConnection) CreateUser(user *User) error {
	insertQuery := fmt.Sprintf(`insert into users (email, first_name, last_name, gender, birthdate, admin_level, password, registration_date)
		values ('%v','%v','%v','%v','%v','%v','%v','%v')`, user.Email, user.Firstname, user.Lastname, user.Gender, user.Birthdate, user.AdminLevel, user.PasswordHash, user.RegistrationDate)
	if results, err := dbc.ExecuteQuery(insertQuery); err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return errors.New("Email already used")
		}
		return err
	} else {
		defer results.Close()
		return nil
	}
}

func (dbc *DatabaseConnection) GetUser(email string) (*User, error) {
	userQuery := fmt.Sprintf("select * from users where email='%v'", email)
	if results, err := dbc.ExecuteQuery(userQuery); err != nil {
		return nil, err
	} else {
		defer results.Close()
		if results.Next() {
			var user User
			if err := results.Scan(
				&user.Email,
				&user.Firstname,
				&user.Lastname,
				&user.AdminLevel,
				&user.About,
				&user.ProfileImageUrl,
				&user.PasswordHash,
				&user.Gender,
				&user.Birthdate,
				&user.RegistrationDate,
				&user.Active,
			); err != nil {
				return nil, err
			}
			if tags, err := dbc.GetTags(email); err != nil {
				return nil, err
			} else {
				user.Tags = tags
				return &user, nil
			}
		} else {
			return nil, nil
		}
	}
}

func (dbc *DatabaseConnection) GetTags(email string) ([]Tag, error) {
	var tags []Tag
	tagsQuery := fmt.Sprintf("select id, content from tags where email='%v'", email)
	if results, err := dbc.ExecuteQuery(tagsQuery); err != nil {
		return nil, err
	} else {
		defer results.Close()
		for results.Next() {
			var tag Tag
			if err := results.Scan(&tag.Id, &tag.Content); err != nil {
				return nil, err
			}
			tags = append(tags, tag)
		}
		return tags, nil
	}
}

func (dbc *DatabaseConnection) ExecuteQuery(query string) (*sql.Rows, error) {
	//db, err := sql.Open("mysql", "capstoneuser:capst0ne18project!@tcp(capstone.cczajq2nppkf.us-east-2.rds.amazonaws.com)/roommates40plus")
	dbstring := fmt.Sprintf("%v:%v@tcp(%v)/%v", dbc.User, dbc.Password, dbc.Endpoint, dbc.DatabaseName)
	db, err := sql.Open("mysql", dbstring)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// Make sure columns match up with user fields
func MapUserToTable(results *sql.Rows, user *User) error {
	return results.Scan(
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.AdminLevel,
		&user.About,
		&user.ProfileImageUrl,
		&user.PasswordHash,
		&user.Gender,
		&user.Birthdate,
		&user.RegistrationDate,
		&user.Active,
	)
}

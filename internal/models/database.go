package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
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
	password := os.Getenv("DB_PASSWORD")
	endpoint := "capstone.cczajq2nppkf.us-east-2.rds.amazonaws.com"
	databaseName := "roommates40plus"
	return &DatabaseConnection{
		User:         user,
		Password:     password,
		Endpoint:     endpoint,
		DatabaseName: databaseName,
	}
}

func (dbc *DatabaseConnection) UserExists(email string) (bool, error) {
	existsQuery := "select exists(select * from users where email=? limit 1);"
	if results, err := dbc.ExecuteQuery(existsQuery, email); err != nil {
		return false, err
	} else {
		defer results.Close()
		results.Next()
		var exists bool
		if err = results.Scan(&exists); err != nil {
			return false, err
		}
		return exists, nil
	}
}

func (dbc *DatabaseConnection) IsActive(email string) (bool, error) {
	activeQuery := "select active from users where email=?"
	if results, err := dbc.ExecuteQuery(activeQuery, email); err != nil {
		return false, err
	} else {
		defer results.Close()
		if results.Next() {
			var active bool
			if err = results.Scan(&active); err != nil {
				return false, err
			}
			return active, nil
		} else {
			return false, errors.New("Account doesn't exist")
		}
	}
}

func (dbc *DatabaseConnection) CreateUser(user *User) error {
	insertQuery := "insert into users (email, first_name, last_name, gender, birthdate, admin_level, password, registration_date) values (?,?,?,?,?,?,?,?)"
	if results, err := dbc.ExecuteQuery(insertQuery, user.Email, user.Firstname, user.Lastname, user.Gender, user.Birthdate, user.AdminLevel, user.PasswordHash, user.RegistrationDate); err != nil {
		if strings.Contains(err.Error(), "Error 1062:") {
			return errors.New("Email already used")
		}
		return err
	} else {
		defer results.Close()
		return nil
	}
}

func (dbc *DatabaseConnection) EditUser(user *User) (*User, error) {
	editQuery := "update users set first_name=?, last_name=?, gender=?, birthdate=?, about=? where email=?"
	if _, err := dbc.ExecuteQuery(editQuery, user.Firstname, user.Lastname, user.Gender, user.Birthdate, user.About, user.Email); err != nil {
		return nil, err
	}
	if err := dbc.SetTags(user.Email, user.Tags); err != nil {
		return nil, err
	}
	return dbc.GetUser(user.Email)
}

func (dbc *DatabaseConnection) GetUser(email string) (*User, error) {
	userQuery := "select * from users where email=?"
	if results, err := dbc.ExecuteQuery(userQuery, email); err != nil {
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
	tagsQuery := "select id, content from tags where email=?"
	if results, err := dbc.ExecuteQuery(tagsQuery, email); err != nil {
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

func (dbc *DatabaseConnection) SetTags(email string, tags []Tag) error {
	deletionQuery := "delete from tags where email=?"
	if _, err := dbc.ExecuteQuery(deletionQuery, email); err != nil {
		return err
	}
	if len(tags) > 0 {
		insertQuery := "insert into tags (email, content) values "
		var tagContent []interface{}
		for _, tag := range tags {
			tagContent = append(tagContent, tag.Content)
			insertQuery += fmt.Sprintf("('%v',?),", email)
		}
		insertQuery = strings.TrimSuffix(insertQuery, ",")
		if _, err := dbc.ExecuteQuery(insertQuery, tagContent...); err != nil {
			return err
		}
	}
	return nil
}

func (dbc *DatabaseConnection) ExecuteQuery(query string, args ...interface{}) (*sql.Rows, error) {
	dbstring := fmt.Sprintf("%v:%v@tcp(%v)/%v", dbc.User, dbc.Password, dbc.Endpoint, dbc.DatabaseName)
	db, err := sql.Open("mysql", dbstring)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return results, nil
}

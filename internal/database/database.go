package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ExecuteQuery(query string) (*sql.Rows, error) {
	db, err := sql.Open("mysql", "capstoneuser:capst0ne18project!@tcp(capstone.cczajq2nppkf.us-east-2.rds.amazonaws.com)/roommates40plus")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	results, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	//fmt.Println(results)
	return results, nil
}

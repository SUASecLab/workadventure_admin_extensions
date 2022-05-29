package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SQLDatabase struct {
	Username string
	Password string
	Hostname string
	Dbname   string
}

func getQueryStringSQL(queryType QueryType) string {
	switch queryType {
	case UserExists:
		return `SELECT COUNT(*) FROM users WHERE uuid = ?`
	case UserIsAdmin:
		return `SELECT COUNT(*) FROM tags WHERE tag="admin" and uuid = ?`
	}
	return ""
}

func (database SQLDatabase) QueryUserInformation(queryType QueryType, uuid string) (bool, error) {
	db, err := sql.Open("mysql", database.Username+":"+database.Password+
		"@("+database.Hostname+":3306)/"+database.Dbname+"?parseTime=true")
	defer db.Close()

	if err != nil {
		return false, err
	}

	err = db.Ping()
	if err != nil {
		return false, err
	}

	var count int

	err = db.QueryRow(getQueryStringSQL(queryType), uuid).Scan(&count)
	if err != nil {
		return false, err
	}

	if count != 1 {
		return false, nil
	}
	return true, nil
}

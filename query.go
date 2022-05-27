package main

import "database/sql"

func queryUserCount(query, uuid string) (bool, error) {
	db, err := sql.Open("mysql", username+":"+password+"@("+hostname+":3306)/"+dbname+"?parseTime=true")
	defer db.Close()

	if err != nil {
		return false, err
	}

	err = db.Ping()
	if err != nil {
		return false, err
	}

	var count int

	err = db.QueryRow(query, uuid).Scan(&count)
	if err != nil {
		return false, err
	}

	if count != 1 {
		return false, nil
	}
	return true, nil
}

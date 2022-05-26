package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserExistsResponse struct {
	Exists bool   `json:"exists"`
	Error  string `json:"error"`
}

func userExistsQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	variables := mux.Vars(r)
	uuidFromRequest := variables["uuid"]

	isValid, errorMessage := isUUIDValid(uuidFromRequest)
	if !isValid {
		w.WriteHeader(403)
		if errorMessage != nil {
			fmt.Fprintf(w, string(errorMessage))
			return
		}
		return
	}

	exists, err := userExists(uuidFromRequest)
	if err != nil {
		response := UserExistsResponse{
			Exists: false,
			Error:  err.Error(),
		}

		responseToString, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			log.Println(marshalErr)
			return
		}
		fmt.Fprintf(w, string(responseToString))
		return
	}

	response := UserExistsResponse{
		Exists: exists,
	}

	responseToString, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, string(responseToString))
}

func userExists(uuid string) (bool, error) {
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
	query := `SELECT COUNT(*) FROM users WHERE uuid = ?`

	err = db.QueryRow(query, uuid).Scan(&count)
	if err != nil {
		return false, err
	}

	if count != 1 {
		return false, nil
	}
	return true, nil
}

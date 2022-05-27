package main

import (
	"encoding/json"
	"log"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func userExistsQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	variables := mux.Vars(r)
	uuidFromRequest := variables["uuid"]

	isValid := extensions.IsUUIDValid(uuidFromRequest)
	if !isValid {
		w.WriteHeader(403)
		response := extensions.UserExistsResponse{
			Exists: false,
			Error:  "The provided UUID is invalid",
		}

		responseToString, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			log.Println(marshalErr)
			return
		}
		fmt.Fprintf(w, string(responseToString))
		return
	}

	exists, err := queryUserCount(`SELECT COUNT(*) FROM users WHERE uuid = ?`, uuidFromRequest)

	if err != nil {
		response := extensions.UserExistsResponse{
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

	response := extensions.UserExistsResponse{
		Exists: exists,
	}

	responseToString, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, string(responseToString))
}

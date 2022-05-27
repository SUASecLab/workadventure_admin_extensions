package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func isAdminQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	variables := mux.Vars(r)
	uuidFromRequest := variables["uuid"]

	isValid := extensions.IsUUIDValid(uuidFromRequest)
	if !isValid {
		w.WriteHeader(403)
		response := extensions.UserIsAdminResponse{
			IsAdmin: false,
			Error:   "The provided UUID is invalid",
		}

		responseToString, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			log.Println(marshalErr)
			return
		}
		fmt.Fprintf(w, string(responseToString))
		return
	}

	// Check if account exists
	exists, err := queryUserCount(`SELECT COUNT(*) FROM users WHERE uuid = ?`, uuidFromRequest)
	if err != nil || !exists {
		response := extensions.UserIsAdminResponse{
			IsAdmin: false,
		}

		if !exists {
			response.Error = "User does not exist"
		}

		if err != nil {
			response.Error = err.Error()
		}

		responseToString, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			log.Println(marshalErr)
			return
		}
		fmt.Fprintf(w, string(responseToString))
		return
	}

	isAdmin, err := queryUserCount(`SELECT COUNT(*) FROM tags WHERE tag="admin" and uuid = ?`, uuidFromRequest)
	if err != nil {
		response := extensions.UserIsAdminResponse{
			IsAdmin: false,
			Error:   err.Error(),
		}

		responseToString, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			log.Println(marshalErr)
			return
		}
		fmt.Fprintf(w, string(responseToString))
		return
	}

	response := extensions.UserIsAdminResponse{
		IsAdmin: isAdmin,
	}

	responseToString, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Fprintf(w, string(responseToString))
}

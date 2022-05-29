package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/SUASecLab/workadventure_admin_extensions/db"
	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

func isAdminHandler(w http.ResponseWriter, r *http.Request) {
	//Send HTTP headers for CORS and enable JSON encoding
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Get UUID from request
	variables := mux.Vars(r)
	uuidFromRequest := variables["uuid"]

	// Create basic response
	response := extensions.UserIsAdminResponse{}

	// Check, if database is defined. If not, we are running a test.
	// Then, use the mock interface.
	if database == nil {
		database = db.MockDatabase{}
	}

	// Check if UUID is valid
	isValid := extensions.IsUUIDValid(uuidFromRequest)
	if !isValid {
		w.WriteHeader(403)
		response.IsAdmin = false
		response.Error = "The provided UUID is invalid"
	} else {
		// Check if user exists
		exists, err := database.QueryUserInformation(db.UserExists, uuidFromRequest)

		if err != nil {
			response.IsAdmin = false
			response.Error = err.Error()
		} else if !exists {
			response.IsAdmin = false
			response.Error = "User does not exist"
		} else {
			// Check if user is admin
			isAdmin, err := database.QueryUserInformation(db.UserIsAdmin, uuidFromRequest)

			if err != nil {
				response.IsAdmin = false
				response.Error = err.Error()
			} else if !isAdmin {
				response.IsAdmin = false
			} else {
				response.IsAdmin = true
			}
		}
	}

	// Encode JSON response
	responseToString, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	// Output response
	fmt.Fprintf(w, string(responseToString))
}

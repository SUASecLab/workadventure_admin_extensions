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

func userExistsHandler(w http.ResponseWriter, r *http.Request) {
	// Send HTTP headers for CORS and enable JSON encoding
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Get UUID from request
	variables := mux.Vars(r)
	uuidFromRequest := variables["uuid"]

	// Create basic response
	response := extensions.UserExistsResponse{}

	// Check, if database is defined. If not, we are running a test.
	// Then, use the mock interface.
	if database == nil {
		database = db.MockDatabase{}
	}

	// Check if UUID is valid
	isValid := extensions.IsUUIDValid(uuidFromRequest)
	if !isValid {
		w.WriteHeader(403)
		response.Exists = false
		response.Error = "The provided UUID is invalid"
	} else {
		// Check if user exists
		exists, err := database.QueryUserInformation(db.UserExists, uuidFromRequest)

		if err != nil {
			response.Exists = false
			response.Error = err.Error()
		} else {
			response.Exists = exists
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

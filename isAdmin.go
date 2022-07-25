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

func isAdminHandler(w http.ResponseWriter, r *http.Request) {
	//Send HTTP headers for CORS and enable JSON encoding
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Validate tags -> if we can decode the token, it is valid.
	// The token is handed out by admin -> if valid, then the tags really exist.
	variables := mux.Vars(r)
	token := variables["token"]

	success, claims := extensions.DecodeToken(token, externalToken)
	if !success {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "The provided token is invalid")
		return
	}

	tags, exists := claims["tags"]
	tagsSlice, okay := tags.([]string)

	if !exists || !okay {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Could not read tags")
		return
	}

	// search for admin tag
	var isAdmin bool
	for _, tag := range tagsSlice {
		if tag == "admin" {
			isAdmin = true
			break
		}
	}

	// Create response
	response := extensions.UserIsAdminResponse{
		IsAdmin: isAdmin,
	}

	// Encode JSON response
	responseToString, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Output response
	fmt.Fprintf(w, string(responseToString))
}

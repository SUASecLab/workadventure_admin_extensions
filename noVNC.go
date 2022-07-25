package main

import (
	"fmt"
	"net/http"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func noVNCPasswordHandler(w http.ResponseWriter, r *http.Request) {
	//Send HTTP headers for CORS and enable JSON encoding
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	// Get token
	token := r.URL.Query().Get("token")

	// Validate token
	success, _ := extensions.DecodeToken(token, externalToken)
	if !success {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "The provided token is invalid")
		return
	}

	// Print password
	fmt.Fprintln(w, string(noVNCPassword))
}

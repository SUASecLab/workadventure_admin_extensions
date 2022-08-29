package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func noVNCPasswordHandler(w http.ResponseWriter, r *http.Request) {
	//Send HTTP headers for CORS and enable JSON encoding
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	// Get token
	token := r.URL.Query().Get("token")

	// Authorize
	noVNCDecision, err := extensions.GetAuthDecision("http://" + sidecarUrl +
		"/auth?token=" + token + "&service=noVNC")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while authorizing")
		log.Println("Could not authorize user:", err)
		return
	}

	if !noVNCDecision.Allowed {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "The provided token is invalid")
		return
	}

	// Print password
	fmt.Fprintln(w, string(noVNCPassword))
}

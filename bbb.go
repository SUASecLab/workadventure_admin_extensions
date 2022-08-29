package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

func bbbApi(apiCall, queryString string) string {
	securityString := apiCall + queryString + bbbSharedSecret
	checksum := sha1.Sum([]byte(securityString))
	stringChecksum := fmt.Sprintf("%x", checksum)
	return queryString + "&checksum=" + stringChecksum
}

func getBBBUrl(w http.ResponseWriter, r *http.Request) {
	//Send HTTP headers for CORS and enable JSON encoding
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	// Get provided information
	token := r.URL.Query().Get("token")
	meetingName := url.QueryEscape(r.URL.Query().Get("meetingName"))
	meetingID := url.QueryEscape(r.URL.Query().Get("meetingID"))
	userName := url.QueryEscape(r.URL.Query().Get("userName"))

	if len(token) == 0 || len(meetingName) == 0 ||
		len(meetingID) == 0 || len(userName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Insufficient parameters provided")
		return
	}

	moderatorDecision, err := extensions.GetAuthDecision("http://" + sidecarUrl +
		"/auth?token=" + token + "&service=bbbModerator")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not authorize user")
		log.Println("Could not authorize user:", err)
		return
	}
	// generate room
	// pasword are secured by being hashed with bbb's secret
	generateRoomUrl := bbbUrl + "api/create?" +
		bbbApi("create", "name="+meetingName+"&meetingID="+meetingID+
			"&attendeePW=attendeePW&moderatorPW=moderatorPW")

	_, err = extensions.Request(generateRoomUrl)

	if err != nil {
		log.Println("Could not generate BBB meeting", err)
	}

	// generate join link
	var generateJoinUrl string

	if moderatorDecision.Allowed {
		generateJoinUrl = bbbUrl + "api/join?" +
			bbbApi("join", "fullName="+userName+"&meetingID="+meetingID+
				"&password=moderatorPW")
	} else {
		generateJoinUrl = bbbUrl + "api/join?" +
			bbbApi("join", "fullName="+userName+"&meetingID="+meetingID+
				"&password=attendeePW")
	}

	fmt.Fprintln(w, generateJoinUrl)
}

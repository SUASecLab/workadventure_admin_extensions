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

	success, userClaims := extensions.DecodeToken(token, externalToken)
	if !success {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "The provided token is invalid")
		return
	}

	isModerator, ok := userClaims["moderator"].(bool)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// generate room
	// pasword are secured by being hashed with bbb's secret
	generateRoomUrl := bbbUrl + "api/create?" +
		bbbApi("create", "name="+meetingName+"&meetingID="+meetingID+
			"&attendeePW=attendeePW&moderatorPW=moderatorPW")

	_, err := extensions.Request(generateRoomUrl)

	if err != nil {
		log.Println("Could not generate BBB meeting", err)
	}

	// generate join link
	var generateJoinUrl string

	if isModerator {
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

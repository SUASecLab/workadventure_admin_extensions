package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
)

type JitsiData struct {
	JitsiUrl string
	RoomName string
	Jwt      string
	UserName string
}

func createJitsiInstance(w http.ResponseWriter, r *http.Request) {
	roomName := url.QueryEscape(r.URL.Query().Get("roomName"))
	userName := url.QueryEscape(r.URL.Query().Get("userName"))
	receivedToken := url.QueryEscape(r.URL.Query().Get("token"))

	issuance, err := extensions.IssueToken("http://" + sidecarUrl + "/issuance?name=" + userName + "&token=" + receivedToken)
	if err != nil {
		log.Println("Could not get Jitsi token: ", err)
		return
	}

	jitsiData := JitsiData{
		JitsiUrl: jitsiUrl,
		RoomName: roomName,
		Jwt:      string(issuance.Token),
		UserName: userName,
	}

	tmpl, err := template.ParseFiles("./templates/jitsi.html")
	if err != nil {
		http.Error(w, "Template file is missing", http.StatusInternalServerError)
		log.Println("Template file not present")
		return
	}
	err = tmpl.Execute(w, jitsiData)
	if err != nil {
		log.Println(err.Error())
	}
}

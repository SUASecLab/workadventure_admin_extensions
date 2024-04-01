package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"text/template"
	"time"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
	"github.com/kataras/jwt"
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

	// does the user have moderator rights?
	moderatorDecision, err := extensions.GetAuthDecision("http://" + sidecarUrl +
		"/auth?token=" + receivedToken + "&service=jitsiModerator")

	// did the query succeed?
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Could not authorize user")
		log.Println("Could not authorize user for Jitsi:", err)
		return
	}

	unixTime := time.Now().Unix()
	jitsiToken, err := jwt.Sign(jwt.HS256, []byte(jitsiKey), map[string]interface{}{
		"context": map[string]interface{}{
			"user": map[string]interface{}{
				"name": userName,
			},
		},
		"nbf":       unixTime - 10,
		"aud":       "jitsi",
		"iss":       jitsiIssuer,
		"room":      "*",
		"moderator": moderatorDecision.Allowed,
		"iat":       unixTime,
		"exp":       unixTime + 60,
	})

	if err != nil {
		fmt.Fprintf(w, "Could not authorize user")
		log.Println("Could not generate Jitsi token", err)
		return
	}

	jitsiData := JitsiData{
		JitsiUrl: jitsiUrl,
		RoomName: roomName,
		Jwt:      string(jitsiToken),
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

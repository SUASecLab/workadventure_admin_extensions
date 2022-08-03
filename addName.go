package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SUASecLab/workadventure_admin_extensions/extensions"
	"github.com/golang-jwt/jwt/v4"
)

func addNameHandler(w http.ResponseWriter, r *http.Request) {
	//Send HTTP headers for CORS and enable JSON encoding
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/plain")

	//Get Name and Token
	name := r.URL.Query().Get("name")
	token := r.URL.Query().Get("token")

	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No name provided")
		return
	}

	success, claims := extensions.DecodeToken(token, externalToken)
	if !success {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "The provided token is invalid")
		return
	}

	// Extend and sign token
	extendedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":      claims["uuid"],
		"tags":      claims["tags"],
		"moderator": claims["moderator"],
		"exp":       claims["exp"],
		"nbf":       claims["nbf"],
		"iat":       claims["iat"],
		"name":      name,
	})

	tokenString, err := extendedToken.SignedString([]byte(externalToken))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("Could not sign token: ", err)
		return
	}

	fmt.Fprintln(w, string(tokenString))
}

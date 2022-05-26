package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	username string
	password string
	hostname string
	dbname   string
)

func main() {
	log.SetFlags(0)

	username = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	hostname = os.Getenv("DB_HOSTNAME")
	dbname = os.Getenv("MYSQL_DATABASE")

	r := mux.NewRouter()
	r.HandleFunc("/userExists/{uuid}", userExistsQuery)
	r.HandleFunc("/isAdmin/{uuid}", isAdminQuery)

	log.Println("Extensions service is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Extensions failed:", err)
	}
}

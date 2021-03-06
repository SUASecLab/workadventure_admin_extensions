package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SUASecLab/workadventure_admin_extensions/db"
	"github.com/gorilla/mux"
)

var (
	database db.Database

	externalToken string
	noVNCPassword string
)

func main() {
	log.SetFlags(0)

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHostname := os.Getenv("DB_HOSTNAME")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	externalToken = os.Getenv("EXTERNAL_TOKEN")
	noVNCPassword = os.Getenv("NOVNC_PASSWORD")

	if len(mysqlUser) == 0 || len(mysqlPassword) == 0 ||
		len(mysqlHostname) == 0 || len(mysqlDatabase) == 0 {
		log.Println("No database information specified. Using mock database for debugging.")
		database = db.MockDatabase{}
	} else {
		database = db.SQLDatabase{
			Username: mysqlUser,
			Password: mysqlPassword,
			Hostname: mysqlHostname,
			Dbname:   mysqlDatabase,
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/userExists/{token}", userExistsHandler)
	r.HandleFunc("/isAdmin/{token}", isAdminHandler)
	r.HandleFunc("/addNameToToken/", addNameHandler)
	r.HandleFunc("/getNoVNCPassword/", noVNCPasswordHandler)

	log.Println("Extensions service is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Extensions failed:", err)
	}
}

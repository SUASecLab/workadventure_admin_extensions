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

	bbbSharedSecret string
	bbbUrl          string
)

func main() {
	log.SetFlags(0)

	externalToken = os.Getenv("EXTERNAL_TOKEN")
	noVNCPassword = os.Getenv("NOVNC_PASSWORD")

	bbbUrl = os.Getenv("BBB_URL")
	bbbSharedSecret = os.Getenv("BBB_SECRET")

	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoHost := os.Getenv("MONGO_HOST")
	mongoDb := os.Getenv("MONGO_DB")

	if len(mongoUsername) == 0 || len(mongoPassword) == 0 ||
		len(mongoHost) == 0 || len(mongoDb) == 0 {
		log.Println("No database information specified. Using mock database for debugging.")
		database = db.MockDatabase{}
	} else {
		database = db.MongoDatabase{
			Username: mongoUsername,
			Password: mongoPassword,
			Hostname: mongoHost,
			Dbname:   mongoDb,
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/userExists/{token}", userExistsHandler)
	r.HandleFunc("/isAdmin/{token}", isAdminHandler)
	r.HandleFunc("/addNameToToken/", addNameHandler)
	r.HandleFunc("/getNoVNCPassword/", noVNCPasswordHandler)
	r.HandleFunc("/bigbluebutton/", getBBBUrl)

	log.Println("Extensions service is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Extensions failed:", err)
	}
}

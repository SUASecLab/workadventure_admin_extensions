package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	sidecarUrl string

	noVNCPassword string

	bbbSharedSecret string
	bbbUrl          string
)

func main() {
	log.SetFlags(0)

	sidecarUrl = os.Getenv("SIDECAR_URL")

	noVNCPassword = os.Getenv("NOVNC_PASSWORD")

	bbbUrl = os.Getenv("BBB_URL")
	bbbSharedSecret = os.Getenv("BBB_SECRET")

	r := mux.NewRouter()
	r.HandleFunc("/getNoVNCPassword/", noVNCPasswordHandler)
	r.HandleFunc("/bigbluebutton/", getBBBUrl)

	log.Println("Extensions service is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Extensions failed:", err)
	}
}

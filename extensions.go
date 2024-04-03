package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	// URL of the sidecar
	sidecarUrl string

	// Password for NoVNC
	noVNCPassword string

	// BBB credentials
	bbbSharedSecret string
	bbbUrl          string

	// Jitsi credentials
	jitsiUrl string

	// Whether to replace BBB by Jitsi
	jitsiReplacesBBB bool

	// subdirectory for extensions service, must start with '/'
	extensionsSubdir string
)

func main() {
	log.SetFlags(0)

	sidecarUrl = os.Getenv("SIDECAR_URL")
	noVNCPassword = os.Getenv("NOVNC_PASSWORD")
	extensionsSubdir = os.Getenv("EXTENSIONS_SUBDIR")

	bbbSharedSecret = os.Getenv("BBB_SECRET")
	bbbUrl = os.Getenv("BBB_URL")

	jitsiUrl = os.Getenv("JITSI_URL")
	jitsiReplacesBBBValue := os.Getenv("JITSI_REPLACES_BBB")
	if jitsiReplacesBBBValue == "true" {
		log.Println("Jitsi replaces BigBlueButton")
		jitsiReplacesBBB = true
	}

	r := mux.NewRouter()
	r.HandleFunc("/getNoVNCPassword/", noVNCPasswordHandler)
	r.HandleFunc("/bigbluebutton/", getBBBUrl)
	r.HandleFunc("/jitsi/", createJitsiInstance)

	log.Println("Extensions service is listening on port 8080")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("Extensions failed:", err)
	}
}

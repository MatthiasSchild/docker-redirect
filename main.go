package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	redirectAddress string
	webAddr         string
)

func init() {
	redirectAddress = os.Getenv("REDIRECT")
	if redirectAddress == "" {
		log.Println("Error: REDIRECT not set.")
		os.Exit(1)
	}

	webAddr = os.Getenv("ADDR")
	if webAddr == "" {
		webAddr = ":8080"
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := strings.TrimSuffix(redirectAddress, "/") + r.RequestURI
		http.Redirect(w, r, url, http.StatusSeeOther)
	})

	log.Println(http.ListenAndServe(webAddr, nil))
}

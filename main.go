package main

import (
	"log"
	"net/http"

	"asciiweb/functions"
)

func main() {
	http.HandleFunc("/", functions.Index)
	http.HandleFunc("/ascii-art", functions.AsciiArt)

	staticDir := "./static/style/"
	staticURL := "/static/style/"
	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle(staticURL, http.StripPrefix(staticURL, fileServer))

	log.Printf("Server started at http://localhost:9000\n")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"asciiweb/printart"
	"asciiweb/read"
)

type Response struct {
	pageTitle string
	// banner    []interface{}
}

const (
	StatusNotFound               = 404 
	StatusBadRequest             = 400
	StatusInternalServerError    = 500
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii-art", asciiArt)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	log.Printf("Server started at http://localhost:9000\n")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		serveError(w, "Page not found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}
	data := Response{
		pageTitle: "ASCII Art Web",
	}
	tmpl.Execute(w, data)
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		serveError(w, "page not found", http.StatusNotFound)
		return
	}
	if strings.ToUpper(r.Method) != http.MethodPost {
		serveError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")

	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		serveError(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// r.Method
	bannerPath := filepath.Join("banner", banner+".txt")
	bannerSlice, err := read.ReadAscii(bannerPath)
	fmt.Println(bannerPath)
	if err != nil {
		serveError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result := printart.AsciiArt(bannerSlice, text)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		serveError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]string{
		"Result": result,
		"Text":   text,
	})
}

func serveError(w http.ResponseWriter, errVal string, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Fatalf("error passing files %v\n", err)
		http.Error(w, "error: %v\n", http.StatusOK)
		return
	}

	code := strconv.Itoa(statusCode)

	w.WriteHeader(statusCode) // Set the HTTP status code

	err = tmpl.Execute(w, struct{ ErrorMsg string }{ErrorMsg: code+" "+errVal})
	if err != nil {
		log.Fatalf("error executing template: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

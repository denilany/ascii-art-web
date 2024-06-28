package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"asciiweb/printart"
	"asciiweb/read"
)

type Response struct {
	pageTitle string
	// banner    []interface{}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/asciiart", asciiArt)
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	log.Printf("Server started at http://localhost:9000\n")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
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
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	if strings.ToUpper(r.Method) != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")

	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// r.Method
	bannerPath := filepath.Join("banner", banner+".txt")
	bannerSlice, err := read.ReadAscii(bannerPath)
	fmt.Println(bannerPath)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	result := printart.AsciiArt(bannerSlice, text)
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]string{
		"Result": result,
		"Text":   text,
	})
}

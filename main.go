package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Response struct {
	pageTitle string
	// banner    []interface{}
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
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

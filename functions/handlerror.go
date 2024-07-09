package functions

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func ServeError(w http.ResponseWriter, errVal string, statusCode int) {
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Fatalf("error passing files %v\n", err)
		http.Error(w, "error: %v\n", http.StatusOK)
		return
	}

	code := strconv.Itoa(statusCode)

	w.WriteHeader(statusCode) // Set the HTTP status code

	err = tmpl.Execute(w, struct{ ErrorMsg string }{ErrorMsg: code + " " + errVal})
	if err != nil {
		log.Fatalf("error executing template: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

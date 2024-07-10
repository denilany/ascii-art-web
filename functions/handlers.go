package functions

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

type Response struct {
	pageTitle string
}

const (
	StatusNotFound            = 404
	StatusBadRequest          = 400
	StatusInternalServerError = 500
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ServeError(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Println("OK: ", http.StatusOK)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data := Response{
		pageTitle: "ASCII Art Web",
	}
	tmpl.Execute(w, data)
}

func Ascii(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		ServeError(w, "page not found", http.StatusNotFound)
		return
	}
	if strings.ToUpper(r.Method) != http.MethodPost {
		ServeError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")

	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		ServeError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	bannerPath := filepath.Join("banner", banner+".txt")
	bannerSlice, err := readAscii(bannerPath)
	fmt.Println(bannerPath)
	if err != nil {
		ServeError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	result, err1 := asciiArt(bannerSlice, text)
	if err1 != nil {
		ServeError(w, "Bad Request", http.StatusBadRequest)
		fmt.Println(err1)

		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ServeError(w, "internal server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, map[string]string{
		"Result": result,
		"Text":   text,
	})
}

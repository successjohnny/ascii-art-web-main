package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Text   string
	Banner string
	Result string
}

// Parse template once when the server starts.
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiHandler)

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "400 Bad Request: Input cannot be empty", http.StatusBadRequest)
		return
	}

	switch banner {
	case "standard", "shadow", "thinkertoy":
		// valid
	default:
		http.Error(w, "400 Bad Request: Invalid banner selected", http.StatusBadRequest)
		return
	}

	result, err := GenerateAscii(text, banner)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			http.Error(w, "404 Not Found: Banner file missing", http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	data := PageData{
		Text:   text,
		Banner: banner,
		Result: result,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
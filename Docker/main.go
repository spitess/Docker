package main

import (
	A "Ascii-Art-Web/asciiart"
	F "fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Result string
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodGet {
		if err := tmpl.Execute(w, Data{Result: ""}); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	if r.Method == http.MethodPost {
		input := r.FormValue("inp1")
		banner := r.FormValue("Files")

		if len(input) > 300 || input == "" {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}

		if !A.InitMap(banner) {
			http.Error(w, " 400 Bad Request", http.StatusBadRequest)
			return
		}

		if err := tmpl.Execute(w, Data{Result: A.Storing(input)}); err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "400 Bad Request", http.StatusBadRequest)
}

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		os.Stderr.WriteString("Err: Invalid Usage [go run .]\n")
		return
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", HandleRequest)

	F.Println("Server running on port http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

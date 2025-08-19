package main

import (
	ftm "fmt"
	"html/template"
	"net/http"
)

func main() {
	// Register the handlers for different routes
	http.HandleFunc("/", index)
	http.HandleFunc("/cv", cv)
	http.HandleFunc("/portfolio", portfolio)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	ftm.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html") // Render the index template
}

func cv(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "cv.html") // Render the index template
}

func portfolio(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "portfolio.html") // Render the index template
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("templates/" + tmpl) // Load the template file
	if err != nil {
		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil) // Execute the template with no data
}

package main

import (
	"net/http"
)

func main() {
	// Serve web files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile("web/prompter.html")
	})
	http.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile("web/privacy.html")
	})
	http.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile("web/app.css")
	})
	http.HandleFunc("/app.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile("web/app.js")
	})

	// API endpoints
	http.Handle("/storescript", scriptStore)
	http.Handle()

	log.Fatal(http.ListenAndServe(":2023", nil))
}
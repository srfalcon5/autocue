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
	http.HandleFunc("/scripts", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("scripts/")
	})
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile("scripts/robots.txt") // Basic privacy for scripts
	})

	// API endpoints
	http.Handle("/storescript", scriptStore)
	http.Handle()

	log.Fatal(http.ListenAndServe(":2023", nil))
}
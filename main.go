package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve web files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/prompter.html")
	})
	http.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/privacy.html")
	})
	http.HandleFunc("/app.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/app.css")
	})
	http.HandleFunc("/app.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/app.js")
	})
	http.HandleFunc("/scripts", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("scripts/"))
	})
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "scripts/robots.txt") // Basic privacy for scripts
	})

	// API endpoints
	http.Handle("/storescript", scriptStore)

	log.Fatal(http.ListenAndServe(":2023", nil))
}
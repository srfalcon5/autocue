package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/robots.txt") // Basic privacy for scripts
	})
	http.HandleFunc("/index.html.br", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/index.html.br")
	})
	http.HandleFunc("/privacy.html.br", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/privacy.html.br")
	})
	http.HandleFunc("/app.css.br", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/app.css.br")
	})

	// API endpoints
	// TODO: rewrite storage script
	http.HandleFunc("/scripts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Allow", "GET, POST")
			// Get script from URL
			code := r.URL.Path[len("/script/"):]
			file := strings.Join([]string{"scripts/", code}, "")

			// Check if script exists
			script, err := os.Open(file)
			if err != nil {
				w.WriteHeader(500)
			} else {
				w.WriteHeader(200)
				w.Header().Add("Prompter-Hash", hash)
				w.Write([]byte(script))
			}
		} else if r.Method == "POST" {
			// Ensure incoming data is of right type
			if r.Header.Get("Content-Type") != "application/json" {
				w.WriteHeader(http.StatusUnsupportedMediaType)
				w.Header().Set("Allow", "GET, POST")
				fmt.Fprintf(w, "Wrong content-type on request.")
				return
			}
			w.Header().Set("Allow", "GET, POST")

			// Generate hash for file
			// Loosely based on https://github.com/cyckl/uploader
			words, err := os.Open("words.json")
			if err != nil {
				w.WriteHeader(500)
				w.Header().Add("Prompter-Hash", "undefined")
				fmt.Fprintf(w, "Couldn't find or open words.json. %v", err)
			}
			var wordBank []string
			err = json.Unmarshal(words, &words)
			if err != nil {
				w.WriteHeader(500)
				w.Header().Add("Prompter-Hash", "undefined")
				fmt.Fprintf(w, "Couldn't unmarshal JSON data: %v", err)
			}
			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(wordBank), func (i, j int) { wordBank[i], wordBank[j] = wordBank[j], wordBank[i] })
			hash := strings.Join([]{ words[0], words[1], words[2]}, "")

			// Store script in a file
			err = ioutil.WriteFile(loc, r.Body, 0644)
			if err != nil {
				w.WriteHeader(500)
				w.Header().Add("Prompter-Hash", "undefined")
				fmt.Fprintf(w, "Couldn't save script as file: %v", err)
			}

			// Return Created with the hash for the client to finish the push
			w.Header().Add("Prompter-Hash", hash)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Allow", "GET, POST")
			fmt.Fprintf(w, "Wrong content-type on request.")
			return
		}
	})

	log.Fatal(http.ListenAndServe(":2023", nil))
}


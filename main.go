package main

import (
	"io/ioutil"
	"log"
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
	http.HandleFunc("/storescript", func(w http.ResponseWriter, r *http.Request) {
		// Parse incoming data
		if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			fmt.Fprintf(w, "Wrong content-type on request.")
			return
		}
		r.ParseForm()
		script := r.FormValue("script")

		// Store in file 
		fname, err := nameGen()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong while pushing to server. Report the error below at https://github.com/doamatto/falcon5-teleprompter/issues/new")
			fmt.Fprintf(w, "%v", err)
			return
		}
		// fname = strings.Join(fname, ".txt") // not sure what this line was for
		f, err := os.Create(fname)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong while saving the your script. Report the error below at https://github.com/doamatto/falcon5-teleprompter/issues/new")
			fmt.Fprintf(w, "%v", err)
			return
		}
		defer f.Close()
		if err := os.WriteFile(fname, []byte(script), 0666); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Something went wrong while saving the your script. Report the error below at https://github.com/doamatto/falcon5-teleprompter/issues/new")
			fmt.Fprintf(w, "%v", err)
			return
		}

		// Strip extension to prevent confusion
		ext := path.Ext(fname)
		hash := fname[0:len(fname)-len(ext)]
		// Return 200 with the hash for the client to finish the push
		w.Header().Add("Prompter-Hash", hash)
		w.WriteHeader(200)
	})
	http.HandleFunc("/script", func(w http.ResponseWriter, r *http.Request) {
		// Get script from URL
		code := r.URL.Path[len("/script/"):]
		file := strings.Join([]string{"scripts/", code, ".txt"}, "")

		// Check if script exists
		script, err := os.Open(file)
		if err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
			w.Header().Add("Prompter-Hash", hash)
			w.Write([]byte(script))
		}
	})

	log.Fatal(http.ListenAndServe(":2023", nil))
}

// Adapted from Daniel's uploader with permission
// https://github.com/cyckl/uploader/blob/master/main.go#L136-L167
func nameGen() (string, error) {
	// Read word file
	data, err := ioutil.ReadFile("./words.json")
	if err != nil {
		return "", errors.New("failed to open word file")
	}
	
	// Link JSON data slice to word slice
	var words []string
	err = json.Unmarshal(data, &words)
	if err != nil {
		return "", errors.New(fmt.Sprintf("could not unmarshal JSON data: %v\n", err))
	}

	// Shuffle words in word list
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	// Get first three entries of shuffled array
	gen := words[0] + words[1] + words[2]
	
	namePre := []string{"./scripts", string(gen), ".txt"}
	name := strings.Join(namePre, "")
	
	return name, nil
}

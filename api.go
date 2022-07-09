package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func scriptStore(w http.ResponseWriter, r *http.Request) {
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
		fmt.FPrintf(w, "%v", err)
		return
	}
	fname := strings.join(fname)
	f, err := os.Create(fname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong while saving the your script. Report the error below at https://github.com/doamatto/falcon5-teleprompter/issues/new")
		fmt.Frptinf(w, err)
		return
	}
	defer f.Close()
	if err := os.WriteFile(fname, []byte(script), 0666); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Something went wrong while saving the your script. Report the error below at https://github.com/doamatto/falcon5-teleprompter/issues/new")
		fmt.Frptinf(w, err)
		return
	}

	// Strip extension to prevent confusion
	ext := path.Ext(fname)
	hash := fname[0:len(fname)-len(ext)]
	// Return 200 with the hash for the client to finish the push
	w.Write(hash)
	w.Header().Add("Prompter-Hash", hash)
	w.WriteHeader(200)
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
	
	name := strings.Join("./scripts/", string(gen), ".txt")
	
	return name, nil
}

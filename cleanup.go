package main

import (
	"log"
	"os"
	"time"
)

func cleanService() {
	// TODO: move this to external automation
	// Cleanup old scripts
	files, err := os.ReadDir("./scripts")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		stat, err := os.Stat(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		mtime := stat.ModTime()
		age := time.Now().Unix() - mtime.Unix()
		if age >= 86400 {
			// 86400s = 1j
			os.Remove(file.Name())
		}
	}
}
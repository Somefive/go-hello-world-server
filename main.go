package main

import (
	"fmt"
	"net/http"
	"os"
)

const VERSION = "0.1.4-rc3"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	username := os.Getenv("USERNAME")
	if username == "" {
		username = "world"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := "USERNAME"
		envs, ok := r.URL.Query()["env"]
		if ok && len(envs) > 0 {
			key = envs[0]
		}
		_, _ = fmt.Fprintf(w, "Hello %s\nVersion: %s\n", os.Getenv(key), VERSION)
	})
	http.HandleFunc("/"+username, func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Welcome\n")
	})
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		println(err.Error())
	}
}

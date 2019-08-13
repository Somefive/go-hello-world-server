package main

import (
	"fmt"
	"net/http"
	"os"
)

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
		fmt.Fprintf(w, "Hello %s\n", os.Getenv(key))
	})
	http.HandleFunc("/"+username, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome\n")
	})
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		println(err.Error())
	}
}

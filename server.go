package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := os.Getenv("NAME")
		age := os.Getenv("AGE")
		msg := fmt.Sprintf("Hello, %s! You are %s years old.", name, age)
		w.Write([]byte(msg))
	})
	http.HandleFunc("/myfamily", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("family/myfamily.txt")
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		w.Write(data)
	})
	http.HandleFunc("/secret", func(w http.ResponseWriter, r *http.Request) {
		user := os.Getenv("USER")
		password := os.Getenv("PASSWORD")
		msg := fmt.Sprintf("Hello, %s! Your password is %s.", user, password)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(msg))
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(startedAt)
		if duration.Seconds() < 10 {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Server has been running for %s, which is too short.", duration)))
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}

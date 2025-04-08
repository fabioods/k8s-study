package main

import (
	"fmt"
	"net/http"
	"os"
)

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
	http.ListenAndServe(":8080", nil)
}

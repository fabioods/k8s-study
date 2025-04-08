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
	http.ListenAndServe(":8080", nil)
}

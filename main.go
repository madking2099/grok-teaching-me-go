package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey Grok, I’m serving!")
	})
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}

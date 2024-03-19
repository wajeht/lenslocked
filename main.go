package main

import (
	"fmt"
	"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "contact")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "home")
	})

	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", nil)
}

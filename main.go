package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleHello)
	fmt.Println("Server starting on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Dutonian! 👋")
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define routes
	http.HandleFunc("/greet", Greet)
	// Starting server
	http.ListenAndServe("localhost:8000", nil)
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

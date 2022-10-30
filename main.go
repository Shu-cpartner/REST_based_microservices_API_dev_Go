package main

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	// Define routes
	// http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	// Starting server
	http.ListenAndServe("localhost:8000", nil)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Abby", "London", "000-7637"},
		{"Bobby", "Paris", "012-7637"},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

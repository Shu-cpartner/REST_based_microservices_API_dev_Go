package app

import (
	"encoding/json"
	"fmt"
	"log"
	"microservicesAPIDevInGolang/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zip_code"`
}

// Privary portのServiceとの依存性を定義
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Abby", "London", "000-7637"},
	// 	{"Bobby", "Paris", "012-7637"},
	// }

	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}

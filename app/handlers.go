package app

import (
	"encoding/json"
	"log"
	"microservicesAPIDevInGolang/service"
	"net/http"
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

package app

import (
	"log"
	"microservicesAPIDevInGolang/domain"
	"microservicesAPIDevInGolang/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodPost)

	// Starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

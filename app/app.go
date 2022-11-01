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
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodPost)

	// Starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

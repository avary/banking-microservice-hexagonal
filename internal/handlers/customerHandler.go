package handlers

import (
	"encoding/json"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type CustomerHandlers struct {
	Service service.CustomerService
	L       *log.Logger
}

// GetAllCustomers is a handler function to get all customers
func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ch.L.Println("Handling GET request on ... /customers")

	customers, err := ch.Service.GetAllCustomers()
	if err != nil {
		http.Error(w, "Unable to retrieve all customers", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(customers); err != nil {
		//http.Error(w, "Unable to encode json", http.StatusBadRequest)
		ch.L.Printf("Error while encoding json : %v", err)
		return
	}

}

// GetCustomerByID is a handler function to get customer by id
func (ch *CustomerHandlers) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	ch.L.Println("Handling GET request on ... /customers/{id}")
	w.Header().Set("Content-Type", "application/json")

	// getting customer id from url path
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["customer_id"])
	ch.L.Printf("id = %v", id)

	customer, err := ch.Service.GetById(id)
	if err != nil {
		ch.writeResponse(w, err.StatusCode, err.Message)
	}
	ch.writeResponse(w, http.StatusOK, customer)
}

func (ch *CustomerHandlers) writeResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		ch.L.Printf("Error while encoding json : %v", err)
		return
	}
}

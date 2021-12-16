package handlers

import (
	"encoding/json"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/errs"
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
	ch.L.Println("Handling GET request on ... /customers")
	w.Header().Set("Content-Type", "application/json")

	customers, err := ch.Service.GetAllCustomers()
	if err != nil {
		ch.writeResponse(w, err, err)
	}
	ch.writeResponse(w, customers, nil)
}

// GetCustomerByID is a handler function to get customer by id
func (ch *CustomerHandlers) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	ch.L.Println("Handling GET request on ... /customers/{id}")
	w.Header().Set("Content-Type", "application/json")

	// getting customer id from url path
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["customer_id"])
	//ch.L.Printf("id = %v", id)

	customer, err := ch.Service.GetById(id)
	if err != nil {
		ch.writeResponse(w, err, err)
	}
	ch.writeResponse(w, customer, nil)
}

func (ch *CustomerHandlers) writeResponse(w http.ResponseWriter, data interface{}, err *errs.AppError) {
	// why? http.StatusOK is already written in a successful response
	if err != nil {
		w.WriteHeader(err.StatusCode)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		ch.L.Printf("Error while encoding json : %v", err)
		return
	}
}

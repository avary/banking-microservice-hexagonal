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
func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	ch.L.Println("Handling GET request on /customers/?status={status}")
	status := r.URL.Query().Get("status")

	customers, err := ch.Service.GetAllCustomers(status)

	if err != nil {
		ch.writeResponse(w, err.StatusCode, err.AsMessage())
	} else {
		ch.writeResponse(w, http.StatusOK, customers)
	}
}

// GetCustomerByID is a handler function to get customer by id
func (ch *CustomerHandlers) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	ch.L.Println("Handling GET request on ... /customers/{id}")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["customer_id"])

	customer, err := ch.Service.GetById(id)
	if err != nil {
		ch.writeResponse(w, err.StatusCode, err.AsMessage())
	} else {
		ch.writeResponse(w, http.StatusOK, customer)
	}
}

func (ch *CustomerHandlers) writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		ch.L.Println("Error encoding response: ", err)
		panic(err)
	}
}

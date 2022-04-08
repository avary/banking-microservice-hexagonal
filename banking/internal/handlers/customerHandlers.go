package handlers

import (
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/service"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/pkg/lib"
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
		lib.RenderJSON(w, err.StatusCode, err.AsMessage())
		ch.L.Println(err.AsMessage())
		return
	} else {
		lib.RenderJSON(w, http.StatusOK, customers)
	}
}

// GetCustomerByID is a handler function to get customer by id
func (ch *CustomerHandlers) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	ch.L.Println("Handling GET request on ... /customers/{id}")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["customer_id"])

	customer, err := ch.Service.GetById(id)
	if err != nil {
		lib.RenderJSON(w, err.StatusCode, err.AsMessage())
	} else {
		lib.RenderJSON(w, http.StatusOK, customer)
	}
}

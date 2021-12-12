package service

import (
	"encoding/json"
	"fmt"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Customers struct {
	l *log.Logger
}

func NewCustomers(l *log.Logger) *Customers {
	return &Customers{l}
}

// GetAllCustomers returns all the products from datastore
func (c *Customers) GetAllCustomers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.l.Println("Handling GET request...")

	// fetch all customers from the datasource
	lc := domain.GetCustomers()

	// serialize/encode the list of products to JSON
	if err := lc.ToJSON(w); err != nil {
		http.Error(w,
			fmt.Sprintf("Unable to encode the list of customers to JSON: %v", err),
			http.StatusBadRequest)
		c.l.Printf("Error while encoding json : %v", err)
		return
	}

	c.l.Printf("Total Customers : %#v", lc.Len())
}

// GetCustomerByID returns a customer by its ID
func (c *Customers) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.l.Println("Handling GET request...")

	// getting id from the url path
	vars := mux.Vars(r)
	id := vars["id"]
	//c.l.Println("Customer id : ", id)

	retrievedCustomer, _, err := domain.GetCustomerById(id)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("Unable to fetch customer with id %s: %v", id, err),
			http.StatusBadRequest)
		c.l.Printf("Error while fetching customer with id %s: %v", id, err)
		return
	}

	// serialize/encode the product to JSON
	if err := json.NewEncoder(w).Encode(retrievedCustomer); err != nil {
		http.Error(w,
			fmt.Sprintf("Unable to encode the customer to JSON: %v", err),
			http.StatusBadRequest)
		c.l.Printf("Error while encoding json : %v", err)
		return
	}

	c.l.Printf("Customer info %v : ", retrievedCustomer)
}

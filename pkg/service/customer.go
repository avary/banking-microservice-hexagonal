package service

import (
	"fmt"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"
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

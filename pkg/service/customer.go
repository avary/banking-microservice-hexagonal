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

// ServeHTTP : Products implicitly implements the http.Handler interface
func (c *Customers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		c.l.Println("Handle GET request")
		c.getAllCustomers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// getAllCustomers returns all the products from datastore
func (c *Customers) getAllCustomers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	c.l.Println("Handling GET request...")

	// fetch all products from the datastore
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

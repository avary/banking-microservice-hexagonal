package handlers

import (
	"encoding/json"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/service"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	Service service.CustomerService
	L       *log.Logger
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ch.L.Println("Handling GET request on ... /customers")

	customers, err := ch.Service.GetAllCustomers()
	if err != nil {
		http.Error(w, "Unable to retrieve all customers", http.StatusBadRequest)
		ch.L.Printf("Error while getting all customers : %v", err)
		return
	}
	if err := json.NewEncoder(w).Encode(customers); err != nil {
		http.Error(w, "Unable to encode json", http.StatusBadRequest)
		ch.L.Printf("Error while encoding json : %v", err)
		return
	}
}

package domain

import "github.com/ashtishad/banking-microservice-hexagonal/internal/errs"

type Customer struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

// CustomerRepository is a SECONDARY PORT on Hexagonal architecture
type CustomerRepository interface {
	// FindAll status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}

package domain

import (
	"github.com/ashtishad/banking-microservice-hexagonal/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/errs"
)

type Customer struct {
	Id          int    `db:"customer_id"`
	Name        string `db:"name"`
	City        string `db:"city"`
	Zipcode     string `db:"zipcode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string `db:"status"`
}

// CustomerRepository is a SECONDARY PORT on Hexagonal architecture
type CustomerRepository interface {
	// FindAll status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(id string) (*Customer, *errs.AppError)
}

// These are the functions that are exposed to the outside world
func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToCustomerResponse() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

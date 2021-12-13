package domain

import (
	"encoding/json"
	"io"
)

// CustomerRepositoryStub is a CustomerRepository implementation
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s *CustomerRepositoryStub) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(s.customers)
}

func (s *CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

// NewCustomerRepositoryStub is a factory method to create a CustomerRepositoryStub
func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	var customersList = []Customer{
		{
			Id:          "1",
			Name:        "John Doe",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "01-01-1990",
			Status:      "active",
		},
		{
			Id:          "2",
			Name:        "Jane Doe",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "01-01-1990",
			Status:      "active",
		},
		{
			Id:          "3",
			Name:        "Con Doe",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1980-01-01",
			Status:      "Active",
		},
		{
			Id:          "4",
			Name:        "Don Doe",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1980-01-01",
			Status:      "Inactive",
		},
		{
			Id:          "5",
			Name:        "Eon Doe",
			City:        "New York",
			Zipcode:     "10001",
			DateOfBirth: "1980-01-01",
			Status:      "Inactive",
		},
	}
	return &CustomerRepositoryStub{customersList}
}

package domain

import (
	"encoding/json"
	"io"
)

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type Customers []*Customer

func (c *Customers) Len() int {
	return len(CustomerList)
}

func GetCustomers() Customers {
	return CustomerList
}

// ToJSON populates the JSON payload from the product struct
func (c *Customers) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(c)
}

var CustomerList = Customers{
	{
		Id:          "1",
		Name:        "Abs Foo",
		City:        "New York",
		Zipcode:     "10001",
		DateOfBirth: "1980-01-01",
		Status:      "Active",
	},
	{
		Id:          "2",
		Name:        "Bon Doe",
		City:        "New York",
		Zipcode:     "10001",
		DateOfBirth: "1980-01-01",
		Status:      "Active",
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

package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// CustomerRepository is a Secondary PORT on Hexagonal architecture
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}

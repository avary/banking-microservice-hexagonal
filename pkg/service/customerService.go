package service

import "github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"

// CustomerService is our PRIMARY PORT
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repoDb domain.CustomerRepoDb
}

func NewCustomerService(repo domain.CustomerRepoDb) DefaultCustomerService {
	return DefaultCustomerService{repoDb: repo}
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	customers, err := s.repoDb.FindAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

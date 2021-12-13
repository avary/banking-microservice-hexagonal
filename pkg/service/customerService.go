package service

import "github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

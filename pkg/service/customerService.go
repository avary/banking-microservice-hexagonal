package service

import (
	"github.com/ashtishad/banking-microservice-hexagonal/internal/errs"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"
)

// CustomerService is our PRIMARY PORT
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetById(id int) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repoDb domain.CustomerRepoDb
}

func NewCustomerService(repo domain.CustomerRepoDb) DefaultCustomerService {
	return DefaultCustomerService{repoDb: repo}
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	customers, err := s.repoDb.FindAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// GetById returns customer by id
func (s DefaultCustomerService) GetById(id int) (*domain.Customer, *errs.AppError) {
	customer, err := s.repoDb.FindById(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

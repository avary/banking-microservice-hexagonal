package service

import (
	"github.com/ashtishad/banking-microservice-hexagonal/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/errs"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"
)

// CustomerService is our PRIMARY PORT
type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetById(id int) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repoDb domain.CustomerRepoDb
}

func NewCustomerService(repo domain.CustomerRepoDb) DefaultCustomerService {
	return DefaultCustomerService{repoDb: repo}
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repoDb.FindAll(status)
	if err != nil || len(customers) == 0 {
		return nil, err
	}

	resp := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		resp = append(resp, c.ToCustomerResponse())
	}
	return resp, err
}

// GetById returns customer by id
func (s DefaultCustomerService) GetById(id int) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repoDb.FindById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToCustomerResponse()

	return &resp, nil
}

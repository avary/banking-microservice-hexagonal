package service

import (
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/domain"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/pkg/errs"
	"strconv"
)

// CustomerService is our PRIMARY PORT
type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetById(id int) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}

// GetAllCustomers returns all customers
func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)
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
	c, err := s.repo.FindById(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	resp := c.ToCustomerResponse()

	return &resp, nil
}

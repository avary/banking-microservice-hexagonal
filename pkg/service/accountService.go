package service

import (
	"github.com/ashtishad/banking-microservice-hexagonal/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/errs"
	"github.com/ashtishad/banking-microservice-hexagonal/pkg/domain"
)

// AccountService is the Primary Port of the Account Service
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.AccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

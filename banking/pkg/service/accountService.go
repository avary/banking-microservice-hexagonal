package service

import (
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/errs"
	domain2 "github.com/ashtishad/banking-microservice-hexagonal/banking/pkg/domain"
)

// AccountService is the Primary Port of the Account Service
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.AccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain2.AccountRepository
}

func NewAccountService(repo domain2.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.AccountResponse, *errs.AppError) {
	err := req.ValidateAccountJSON()
	if err != nil {
		return nil, err
	}
	a := domain2.NewAccount(req.CustomerId, req.AccountType, req.Amount)

	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	// incoming request validation
	err := req.ValidateTransactionJSON()
	if err != nil {
		return nil, err
	}

	// fetching the account
	account, err := s.repo.FindById(req.AccountId, req.CustomerId)
	if err != nil {
		return nil, errs.NewNotFoundError("Could not find account")
	}

	// server side validation for checking the available balance in the account
	if req.IsTransactionTypeWithdrawal() {
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	// if all is well, build the transaction object & save the transaction
	t := domain2.NewTransaction(req)

	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

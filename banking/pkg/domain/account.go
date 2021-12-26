package domain

import (
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/errs"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/internal/lib"
	"time"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	GetAccount(accountId string) (*Account, *errs.AppError)
	FindById(accountId string, customerId string) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func NewAccount(customerId, accountType string, amount float64) Account {
	return Account{
		CustomerId:  customerId,
		OpeningDate: time.Now().Format(lib.DbTSLayout),
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

func (a Account) ToNewAccountResponseDto() dto.AccountResponse {
	return dto.AccountResponse{AccountId: a.AccountId}
}

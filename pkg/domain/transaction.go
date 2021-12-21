package domain

import (
	"github.com/ashtishad/banking-microservice-hexagonal/internal/dto"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/lib"
	"time"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionDate string  `db:"transaction_date"`
}

func NewTransaction(transactionDto dto.TransactionRequest) Transaction {
	return Transaction{
		AccountId:       transactionDto.AccountId,
		Amount:          transactionDto.Amount,
		TransactionType: transactionDto.TransactionType,
		TransactionDate: time.Now().Format(lib.DbTSLayout),
	}
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == lib.WITHDRAWAL {
		return true
	}
	return false
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

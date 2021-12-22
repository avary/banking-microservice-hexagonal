package domain

import (
	"database/sql"
	"fmt"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/errs"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"strconv"
)

type AccountRepoDb struct {
	db *sql.DB
	L  *log.Logger
}

// NewAccountRepoDb creates a new AccountRepoDb wth a *sql.DB
func NewAccountRepoDb(dbClient *sql.DB, L *log.Logger) AccountRepoDb {
	return AccountRepoDb{dbClient, L}
}

// Save inserts an account into the database
func (d AccountRepoDb) Save(a Account) (*Account, *errs.AppError) {
	d.L.Println("Inserting a:", a)

	// check customer id exists
	if err := CheckCustomerIdExists(d.db, a.CustomerId); err != nil {
		d.L.Printf("Customer not found in database, can't create account : %s", err.Error())
		return nil, errs.NewNotFoundError("Customer not found in database, can't create account")
	}

	// create account for this existing customer on postgres
	sqlStatement := `
	INSERT INTO accounts (customer_id, opening_date, account_type, amount,status)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING account_id`
	var id int64
	err := d.db.QueryRow(sqlStatement, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New account ID is:", id)

	// set the id of the account. why strconv and why not string(id)? because the id is a string,
	// normally string(id) converts it to unicode codepoint
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

// SaveTransaction inserts a transaction into the database, rollback if error
func (d AccountRepoDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// starting the database transaction block
	tx, err := d.db.Begin()
	if err != nil {
		d.L.Printf("Error starting transaction : %s", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	sqlStatement := `INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
											values ($1, $2, $3, $4) RETURNING transaction_id`

	var transactionId int64
	err = d.db.QueryRow(sqlStatement, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate).Scan(&transactionId)
	if err != nil {
		d.L.Printf("Error inserting transaction : %s", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - $1 where account_id = $2`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + $1 where account_id = $2`, t.Amount, t.AccountId)
	}

	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		_ = tx.Rollback()
		d.L.Printf("Error while saving transaction : %s", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		d.L.Printf("Error while committing transaction : %s", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Getting the latest account information from the accounts table
	account, appErr := d.FindById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepoDb) FindById(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = $1"
	row := d.db.QueryRow(sqlGetAccount, accountId)

	var a Account
	err := row.Scan(&a.AccountId, &a.CustomerId, &a.OpeningDate, &a.AccountType, &a.Amount)
	if err == sql.ErrNoRows {
		d.L.Printf("Error while scanning accounts by id : %s", err.Error())
		return nil, errs.NewNotFoundError("Account not found")
	}
	if err != nil {
		d.L.Printf("Error while getting account by id : %s", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &a, nil
}

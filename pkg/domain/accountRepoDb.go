package domain

import (
	"database/sql"
	"github.com/ashtishad/banking-microservice-hexagonal/internal/errs"
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
func (r AccountRepoDb) Save(a Account) (*Account, *errs.AppError) {
	r.L.Println("Inserting a:", a)

	// check customer id exists
	if err := CheckCustomerIdExists(r.db, a.CustomerId); err != nil {
		r.L.Printf("Customer not found in database, can't create account : %s", err.Error())
		return nil, errs.NewNotFoundError("Customer not found in database, can't create account")
	}

	// create account for this existing customer
	stmt, err := r.db.Prepare("INSERT INTO accounts (customer_id, opening_date, account_type, amount,status) VALUES (?, ?, ?, ? , ?)")
	if err != nil {
		r.L.Printf("Error while preparing statement : %s", err.Error())
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer statementClose(stmt)

	res, err := stmt.Exec(a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		r.L.Printf("Error while executing statement : %s", err.Error())
		return nil, errs.NewInternalServerError(err.Error())
	}

	// get the id of the account that was just inserted
	id, err := res.LastInsertId()
	if err != nil {
		r.L.Printf("Error while getting last insert id : %s", err.Error())
		return nil, errs.NewInternalServerError(err.Error())
	}
	// set the id of the account. why strconv and why not string(id)? because the id is a string,
	// normally string(id) converts it to unicode codepoint
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func CheckCustomerIdExists(db *sql.DB, reqCustomerId string) error {
	var c Customer
	findIdSql := "select customer_id from customers where customer_id = ?"
	row := db.QueryRow(findIdSql, reqCustomerId)
	err := row.Scan(&c.Id)
	if err == sql.ErrNoRows {
		return err
	}
	return nil
}

func statementClose(stmt *sql.Stmt) {
	err := stmt.Close()
	if err != nil {
		log.Printf("Error while closing the statement: %v\n", err)
	}
}

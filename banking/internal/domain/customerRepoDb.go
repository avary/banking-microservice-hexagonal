package domain

import (
	"database/sql"
	"github.com/ashtishad/banking-microservice-hexagonal/banking/pkg/errs"
	"log"
)

type CustomerRepoDb struct {
	db *sql.DB
	L  *log.Logger
}

// NewCustomerRepoDb creates a new customer repository
func NewCustomerRepoDb(dbClient *sql.DB, L *log.Logger) CustomerRepoDb {
	return CustomerRepoDb{dbClient, L}
}

// FindAll returns all customers from the database
func (d CustomerRepoDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers`
		rows, err = d.db.Query(findAllSql)
	} else if status == "1" || status == "0" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = $1"
		rows, err = d.db.Query(findAllSql, status)
	} else {
		return nil, errs.NewNotFoundError("status is not valid")
	}

	if err != nil {
		d.L.Println("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			d.L.Printf("Error while scanning customers :  %v", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// FindById returns a customer by id
func (d CustomerRepoDb) FindById(id string) (*Customer, *errs.AppError) {
	// Note: Select * would supply data on db table order, order would mismatch with struct fields
	findByIdSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = $1"
	row := d.db.QueryRow(findByIdSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err == sql.ErrNoRows {
		d.L.Printf("Error while scanning customers by id : %s", err.Error())
		return nil, errs.NewNotFoundError("Customer not found in database")
	}
	// catch other errors that might occur
	if err != nil {
		d.L.Printf("Error while scanning customers by id : %s", err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

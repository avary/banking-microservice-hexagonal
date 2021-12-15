package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepoDb struct {
	db *sql.DB
	L  *log.Logger
}

// NewCustomerRepoDb creates a new customer repository
// https://github.com/go-sql-driver/mysql
func NewCustomerRepoDb(L *log.Logger) CustomerRepoDb {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepoDb{db, L}
}

// FindAll returns all customers from the database
func (d *CustomerRepoDb) FindAll() ([]Customer, error) {
	findAllSql := "select * from customers"
	rows, err := d.db.Query(findAllSql)
	if err != nil {
		d.L.Printf("Error while querying on customer table : %s", err.Error())
	}
	// defer rows.Close() , error wrapped in a closure
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			d.L.Printf("Error while closing rows : %s", err.Error())
		}
	}(rows)

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status); err != nil {
			d.L.Printf("Error while scanning customers table : %s", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

// FindById returns a customer by id
func (d *CustomerRepoDb) FindById(id int) (Customer, error) {
	findByIdSql := "select * from customers where customer_id = ?"
	row := d.db.QueryRow(findByIdSql, id)

	var c Customer
	if err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status); err != nil {
		d.L.Printf("Error while scanning customers by id : %s", err.Error())
		return c, err
	}
	return c, nil
}

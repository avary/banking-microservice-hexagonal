package domain

import "database/sql"

func CheckCustomerIdExists(db *sql.DB, reqCustomerId string) error {
	var c Customer
	findIdSql := "select customer_id from customers where customer_id = $1"
	row := db.QueryRow(findIdSql, reqCustomerId)
	err := row.Scan(&c.Id)
	if err == sql.ErrNoRows {
		return err
	}
	return nil
}

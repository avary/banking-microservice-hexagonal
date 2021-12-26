package domain

import (
	"database/sql"
	"errors"
	"log"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, error)
}

type AuthRepositoryDb struct {
	client *sql.DB
}

func NewAuthRepository(client *sql.DB) AuthRepositoryDb {
	return AuthRepositoryDb{client}
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, error) {
	// TODO: Improve sql statement
	sqlVerify := `SELECT username, u.customer_id, role, array_to_string(array_agg(a.account_id), ',') as account_numbers 
				  FROM users u
                  LEFT JOIN accounts a ON a.customer_id = u.customer_id
                  WHERE username = $1 and password = $2
                  GROUP BY u.customer_id , username, role`

	row := d.client.QueryRow(sqlVerify, username, password)

	var login Login
	err := row.Scan(&login.Username, &login.CustomerId, &login.Role, &login.Accounts)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		} else {
			log.Println("Error while verifying login request from database: " + err.Error())
			return nil, errors.New("unexpected database error")
		}
	}
	return &login, nil
}

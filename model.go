// model.go

package main

import (
	"database/sql"
)

type account struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Api_token string `json:"api_token"`
}

func (a *account) createAccount(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO account(name, email, api_token) VALUES($1, $2, $3) RETURNING id",
		a.Name, a.Email, a.Api_token).Scan(&a.ID)

	if err != nil {
		return err
	}

	return nil
}

func (a *account) getAccount(db *sql.DB) error {
	return db.QueryRow("SELECT name, email, api_token FROM account WHERE id=$1",
		a.ID).Scan(&a.Name, &a.Email, &a.Api_token)
}

func (a *account) updateAccount(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE account SET name=$1, email=$2, api_token=$3 WHERE id=$4",
			a.Name, a.Email, a.Api_token, a.ID)

	return err
}

func (a *account) deleteAccount(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM account WHERE id=$1", a.ID)

	return err
}

func listAccounts(db *sql.DB) ([]account, error) {
	rows, err := db.Query("SELECT * FROM account")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []account{}

	for rows.Next() {
		var a account
		if err := rows.Scan(&a.ID, &a.Name, &a.Email, &a.Api_token); err != nil {
			return nil, err
		}
		accounts = append(accounts, a)
	}

	return accounts, nil
}

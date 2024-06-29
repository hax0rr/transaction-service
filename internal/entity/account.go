package entity

import "fmt"

type Account struct {
	ID             string `db:"id"`
	DocumentNumber string `db:"document_number"`
}

func (acc *Account) Validate() error {
	if acc == nil || len(acc.DocumentNumber) == 0 {
		return fmt.Errorf("document_number can't be empty")
	}

	return nil
}

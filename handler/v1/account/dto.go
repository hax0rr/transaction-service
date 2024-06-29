package account

import "errors"

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number,omitempty'"`
}

type CreateAccountResponse struct {
	ID string `json:"account_id"`
}

type GetAccountByIDResponse struct {
	ID             string `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (req *CreateAccountRequest) Validate() error {
	if req == nil || len(req.DocumentNumber) == 0 {
		return errors.New("document_number can't be empty")
	}

	return nil
}

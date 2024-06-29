package transaction

import (
	"errors"
)

type CreateTransactionRequest struct {
	AccountID       string  `json:"account_id,omitempty"`
	OperationTypeID int     `json:"operation_type_id,omitempty"`
	Amount          float64 `json:"amount,omitempty"`
}

type CreateTransactionResponse struct {
	ID string `json:"transaction_id"`
}

func (req *CreateTransactionRequest) Validate() error {
	if req == nil {
		return errors.New("document_number can't be empty")
	}

	if len(req.AccountID) == 0 {
		return errors.New("account_id can't be empty")
	}

	return nil
}

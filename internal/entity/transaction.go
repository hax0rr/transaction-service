package entity

import "fmt"

type Transaction struct {
	ID              string  `db:"id"`
	AccountID       string  `db:"account_id"`
	OperationTypeID int     `db:"operation_type_id"`
	Amount          float64 `db:"amount"`
}

func (txn *Transaction) Validate() error {
	opType := OperationType(txn.OperationTypeID)
	if !opType.IsValid() {
		return fmt.Errorf("invalid operation_type_id %d", txn.OperationTypeID)
	}

	if txn.Amount <= 0 {
		return fmt.Errorf("invalid amount %f", txn.Amount)
	}

	return nil
}

func (txn *Transaction) SetTxnAmount() {
	opType := OperationType(txn.OperationTypeID)
	txn.Amount = opType.GetAmount(txn.Amount)
}

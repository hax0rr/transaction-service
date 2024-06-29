package entity

type OperationType int

const (
	NormalPurchase          OperationType = 1
	PurchaseWithInstallment OperationType = 2
	Withdrawal              OperationType = 3
	CreditVoucher           OperationType = 4
)

func (op *OperationType) IsValid() bool {
	switch *op {
	case NormalPurchase, PurchaseWithInstallment, Withdrawal, CreditVoucher:
		return true
	}
	return false
}

func (op *OperationType) GetAmount(amount float64) float64 {
	switch *op {
	case NormalPurchase, PurchaseWithInstallment, Withdrawal:
		return amount * -1
	case CreditVoucher:
		return amount
	}

	return 0
}

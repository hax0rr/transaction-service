package entity

import (
	"testing"
)

func TestTransaction_SetTxnAmount(t *testing.T) {
	type fields struct {
		ID              string
		AccountID       string
		OperationTypeID int
		Amount          float64
		WantAmount      float64
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "positive amount",
			fields: fields{
				OperationTypeID: 4,
				Amount:          22,
				WantAmount:      22,
			},
		},
		{
			name: "negative amount",
			fields: fields{
				OperationTypeID: 1,
				Amount:          22,
				WantAmount:      -22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txn := &Transaction{
				ID:              tt.fields.ID,
				AccountID:       tt.fields.AccountID,
				OperationTypeID: tt.fields.OperationTypeID,
				Amount:          tt.fields.Amount,
			}
			txn.SetTxnAmount()
			if txn.Amount != tt.fields.WantAmount {
				t.Errorf("GetAmount() = %v, want %v", txn.Amount, tt.fields.WantAmount)
			}
		})
	}
}

func TestTransaction_Validate(t *testing.T) {
	type fields struct {
		ID              string
		AccountID       string
		OperationTypeID int
		Amount          float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				OperationTypeID: 1,
				Amount:          1,
			},
			wantErr: false,
		},
		{
			name: "invalid-1",
			fields: fields{
				OperationTypeID: 5,
				Amount:          1,
			},
			wantErr: true,
		},
		{
			name: "invalid-2",
			fields: fields{
				OperationTypeID: 1,
				Amount:          0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			txn := &Transaction{
				ID:              tt.fields.ID,
				AccountID:       tt.fields.AccountID,
				OperationTypeID: tt.fields.OperationTypeID,
				Amount:          tt.fields.Amount,
			}
			if err := txn.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

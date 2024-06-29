package entity

import "testing"

func TestOperationType_GetAmount(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		op   OperationType
		args args
		want float64
	}{
		{
			name: "success negative amount for purchase",
			op:   NormalPurchase,
			args: args{amount: 1},
			want: -1,
		},
		{
			name: "success positive amount for purchase",
			op:   CreditVoucher,
			args: args{amount: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.GetAmount(tt.args.amount); got != tt.want {
				t.Errorf("GetAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		op   OperationType
		want bool
	}{
		{
			name: "invalid",
			op:   12,
			want: false,
		},
		{
			name: "valid",
			op:   CreditVoucher,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

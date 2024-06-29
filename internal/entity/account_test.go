package entity

import "testing"

func TestAccount_Validate(t *testing.T) {
	type fields struct {
		ID             string
		DocumentNumber string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ID:             "1",
				DocumentNumber: "1",
			},
			wantErr: false,
		},
		{
			name:    "error",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc := &Account{
				ID:             tt.fields.ID,
				DocumentNumber: tt.fields.DocumentNumber,
			}
			if err := acc.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

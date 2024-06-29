package transaction

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"transaction-service/internal/entity"
	"transaction-service/internal/repository"
)

func TestService_CreateTransaction(t *testing.T) {
	type fields struct {
		repository func(controller *gomock.Controller) repository.IRepository
	}
	type args struct {
		ctx         context.Context
		transaction *entity.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Transaction
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repository: func(controller *gomock.Controller) repository.IRepository {
					mr := repository.NewMockIRepository(controller)
					mr.EXPECT().
						CreateTransaction(gomock.Any(), gomock.Any()).
						Times(1).
						Return(&entity.Transaction{
							ID:              "11",
							AccountID:       "123",
							OperationTypeID: 1,
							Amount:          -23,
						}, nil)
					return mr
				},
			},
			args: args{
				ctx: context.TODO(),
				transaction: &entity.Transaction{
					AccountID:       "123",
					OperationTypeID: 1,
					Amount:          23,
				},
			},
			want: &entity.Transaction{
				ID:              "11",
				AccountID:       "123",
				OperationTypeID: 1,
				Amount:          -23,
			},
			wantErr: false,
		},
		{
			name: "repo returns error",
			fields: fields{
				repository: func(controller *gomock.Controller) repository.IRepository {
					mr := repository.NewMockIRepository(controller)
					mr.EXPECT().
						CreateTransaction(gomock.Any(), gomock.Any()).
						Times(1).
						Return(nil, errors.New("some error"))
					return mr
				},
			},
			args: args{
				ctx: context.TODO(),
				transaction: &entity.Transaction{
					AccountID:       "123",
					OperationTypeID: 1,
					Amount:          23,
				},
			},

			wantErr: true,
		},
		{
			name: "validation fails",
			fields: fields{
				repository: func(controller *gomock.Controller) repository.IRepository {
					mr := repository.NewMockIRepository(controller)
					return mr
				},
			},
			args: args{
				ctx: context.TODO(),
				transaction: &entity.Transaction{
					AccountID:       "123",
					OperationTypeID: 11,
					Amount:          23,
				},
			},

			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			s := New(tt.fields.repository(controller))
			got, err := s.CreateTransaction(tt.args.ctx, tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

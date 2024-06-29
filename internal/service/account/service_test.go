package account

import (
	"context"
	"errors"
	"github.com/hax0rr/transaction-service/internal/entity"
	"github.com/hax0rr/transaction-service/internal/repository"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestService_CreateAccount(t *testing.T) {
	type fields struct {
		repository func(ctrl *gomock.Controller) repository.IRepository
	}
	type args struct {
		ctx     context.Context
		account *entity.Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Account
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:     context.TODO(),
				account: &entity.Account{DocumentNumber: "1"},
			},
			fields: fields{
				repository: func(ctrl *gomock.Controller) repository.IRepository {
					rm := repository.NewMockIRepository(ctrl)
					rm.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).
						Return(&entity.Account{DocumentNumber: "1", ID: "123"}, nil).
						Times(1)

					return rm
				},
			},
			wantErr: false,
			want: &entity.Account{
				ID:             "123",
				DocumentNumber: "1",
			},
		},
		{
			name: "repository returns some error",
			args: args{
				ctx:     context.TODO(),
				account: &entity.Account{DocumentNumber: "1"},
			},
			fields: fields{
				repository: func(ctrl *gomock.Controller) repository.IRepository {
					mr := repository.NewMockIRepository(ctrl)
					mr.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).
						Return(nil, errors.New("error")).
						Times(1)

					return mr
				},
			},
			wantErr: true,
		},
		{
			name: "service validation fails",
			args: args{
				ctx:     context.TODO(),
				account: &entity.Account{DocumentNumber: ""},
			},
			fields: fields{
				repository: func(ctrl *gomock.Controller) repository.IRepository {
					rm := repository.NewMockIRepository(ctrl)
					rm.EXPECT().CreateAccount(gomock.Any(), gomock.Any()).
						Return(nil, errors.New("error")).
						Times(0)

					return rm
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			s := New(tt.fields.repository(controller))
			got, err := s.CreateAccount(tt.args.ctx, tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetAccountByID(t *testing.T) {
	type fields struct {
		repository func(controller *gomock.Controller) repository.IRepository
	}
	type args struct {
		ctx       context.Context
		accountID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Account
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repository: func(controller *gomock.Controller) repository.IRepository {
					mr := repository.NewMockIRepository(controller)
					mr.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Return(&entity.Account{ID: "1", DocumentNumber: "11"}, nil).Times(1)

					return mr
				},
			},
			args: args{
				ctx:       context.TODO(),
				accountID: "1",
			},
			wantErr: false,
			want: &entity.Account{
				ID:             "1",
				DocumentNumber: "11",
			},
		},
		{
			name: "repo fails",
			fields: fields{
				repository: func(controller *gomock.Controller) repository.IRepository {
					mr := repository.NewMockIRepository(controller)
					mr.EXPECT().GetAccountByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("some error")).Times(1)

					return mr
				},
			},
			args: args{
				ctx:       context.TODO(),
				accountID: "1",
			},
			wantErr: true,
		},
		{
			name: "service validation fails",
			fields: fields{
				repository: func(controller *gomock.Controller) repository.IRepository {
					mr := repository.NewMockIRepository(controller)

					return mr
				},
			},
			args: args{
				ctx:       context.TODO(),
				accountID: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			s := New(tt.fields.repository(controller))
			got, err := s.GetAccountByID(tt.args.ctx, tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

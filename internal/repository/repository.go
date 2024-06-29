package repository

import (
	"context"
	"github.com/hax0rr/transaction-service/internal/entity"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	createAccount     = `INSERT into accounts (id, document_number) VALUES (:id, :document_number)`
	getAccountByID    = `select id, document_number from accounts where id=$1`
	createTransaction = `INSERT into transactions (id, account_id, operation_type_id, amount) VALUES (:id, :account_id, :operation_type_id, :amount)`
)

type IRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error)
	GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error)
	CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error)
}

type Repository struct {
	db            *sqlx.DB
	dbTimeoutInMs time.Duration
}

func New(db *sqlx.DB, timeoutInMs int) *Repository {
	return &Repository{db: db, dbTimeoutInMs: time.Millisecond * time.Duration(timeoutInMs)}
}

func (r *Repository) GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error) {
	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeoutInMs)
	defer cancel()

	var account entity.Account

	err := r.db.GetContext(dbCtx, &account, getAccountByID, accountID)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *Repository) CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeoutInMs)
	defer cancel()

	_, err := r.db.NamedExecContext(dbCtx, createAccount, &account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *Repository) CreateTransaction(ctx context.Context, txn *entity.Transaction) (*entity.Transaction, error) {
	dbCtx, cancel := context.WithTimeout(ctx, r.dbTimeoutInMs)
	defer cancel()

	_, err := r.db.NamedExecContext(dbCtx, createTransaction, txn)
	if err != nil {
		return nil, err
	}
	return txn, nil
}

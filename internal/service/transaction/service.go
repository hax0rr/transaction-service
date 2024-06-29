package transaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/hax0rr/transaction-service/internal/entity"
	"github.com/hax0rr/transaction-service/internal/repository"
)

var ErrValidationFailed = errors.New("validation failed")

type IService interface {
	CreateTransaction(ctx context.Context, txn *entity.Transaction) (*entity.Transaction, error)
}

type Service struct {
	repository repository.IRepository
}

func New(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (*entity.Transaction, error) {
	if err := transaction.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrValidationFailed, err)
	}

	transaction.SetTxnAmount()

	res, err := s.repository.CreateTransaction(ctx, transaction)
	if err != nil {
		return nil, err
	}

	return res, nil
}

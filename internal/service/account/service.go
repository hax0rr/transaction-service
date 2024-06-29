package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/hax0rr/transaction-service/internal/entity"
	"github.com/hax0rr/transaction-service/internal/repository"
)

var ErrAccountNotFound = errors.New("account not found")
var ErrValidationFailed = errors.New("validation failed")

type IService interface {
	CreateAccount(ctx context.Context, request *entity.Account) (*entity.Account, error)
	GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error)
}

type Service struct {
	repository repository.IRepository
}

func New(repository repository.IRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAccountByID(ctx context.Context, accountID string) (*entity.Account, error) {
	if len(accountID) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrValidationFailed, "account id can't be empty")
	}

	res, err := s.repository.GetAccountByID(ctx, accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, err
	}

	return res, nil
}

func (s *Service) CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	if err := account.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrValidationFailed, err)
	}

	res, err := s.repository.CreateAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	return res, nil
}

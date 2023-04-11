package account

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type AccountRepository interface {
	FindById(ctx context.Context, id int) (domain.Account, error)
	UpdateBalance(ctx context.Context, account domain.Account) error
	Create(ctx context.Context, account domain.Account) (domain.Account, error)
}

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(
	repo AccountRepository,
) AccountService {
	return AccountService{repo}
}

func (s AccountService) UpdateBalance(ctx context.Context, id int, amount float64) error {
	account, err := s.repo.FindById(ctx, id)
	if err != nil {
		return err
	}

	account.Balance += amount

	return s.repo.UpdateBalance(ctx, account)
}

func (s AccountService) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	return s.repo.Create(ctx, account)
}

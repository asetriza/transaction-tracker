package account

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type AccountRepository interface {
	FindById(id int) (domain.Account, error)
	UpdateBalance(account domain.Account) error
	CreateAccount(ctx context.Context, account domain.Account) (domain.Account, error)
}

type Account struct {
	repo AccountRepository
}

func NewAccount(repo AccountRepository) Account {
	return Account{repo}
}

func (s Account) UpdateBalance(id int, amount float64) error {
	account, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	account.Balance += amount

	return s.repo.UpdateBalance(account)
}

func (s Account) CreateAccount(ctx context.Context, account domain.Account) (domain.Account, error) {
	return s.repo.CreateAccount(ctx, account)
}

package account

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type Service interface {
	UpdateBalance(id int, amount float64) error
	CreateAccount(ctx context.Context, account domain.Account) (domain.Account, error)
}

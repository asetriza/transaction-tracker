package account

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type Service interface {
	UpdateBalance(ctx context.Context, id int, amount float64) error
	Create(ctx context.Context, account domain.Account) (domain.Account, error)
}

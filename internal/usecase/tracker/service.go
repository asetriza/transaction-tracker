package tracker

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type Service interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
}

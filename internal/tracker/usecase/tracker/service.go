package tracker

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/tracker/domain"
)

type Service interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
}

package transaction

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type Service interface {
	Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	CancelLatestOddRecords(ctx context.Context, limit int) error
}

package tracker

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type Repository interface {
	CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
}

type tracker struct {
	repo Repository
}

func NewTracker(repo Repository) tracker {
	return tracker{
		repo: repo,
	}
}

func (t tracker) CreateTransaction(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	if err := transaction.Validate(); err != nil {
		return domain.Transaction{}, err
	}

	transaction, err := t.repo.CreateTransaction(ctx, transaction)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

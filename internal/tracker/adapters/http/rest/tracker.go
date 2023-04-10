package rest

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/tracker/domain"
	models "github.com/asetriza/transaction-tracker/internal/tracker/models"
)

func (h Handler) CreateTransaction(ctx context.Context, req *models.Transaction, params models.CreateTransactionParams) (models.CreateTransactionRes, error) {
	dTransaction, err := toDomainTransaction(*req)
	if err != nil {
		return nil, err
	}

	transaction, err := h.tracker.CreateTransaction(ctx, dTransaction)
	if err != nil {
		return &models.Transaction{}, err
	}

	return fromDomainTransaction(transaction), nil
}

func toDomainTransaction(tr models.Transaction) (domain.Transaction, error) {
	state, err := domain.StringToState(string(tr.GetState()))
	if err != nil {
		return domain.Transaction{}, err
	}

	return domain.Transaction{
		State:         state,
		Amount:        tr.Amount,
		TransactionID: domain.ID(tr.TransactionId),
	}, nil
}

func fromDomainTransaction(tr domain.Transaction) *models.Transaction {
	return &models.Transaction{
		State:         models.TransactionState(tr.State.String()),
		Amount:        tr.Amount,
		TransactionId: string(tr.TransactionID),
	}
}

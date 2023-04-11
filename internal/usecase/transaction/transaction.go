package transaction

import (
	"context"
	"log"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	FindLatestOddRecords(ctx context.Context, limit int) ([]domain.Transaction, error)
	MarkAsCanceled(ctx context.Context, transaction domain.Transaction) error
}

type AccountService interface {
	UpdateBalance(ctx context.Context, id int, amount float64) error
	Create(ctx context.Context, account domain.Account) (domain.Account, error)
}

type TransactionService struct {
	trRepo    TransactionRepository
	acService AccountService
}

func NewTransactionService(
	trRepo TransactionRepository,
	acService AccountService,
) TransactionService {
	return TransactionService{
		trRepo:    trRepo,
		acService: acService,
	}
}

func (t TransactionService) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	if err := transaction.Validate(); err != nil {
		return domain.Transaction{}, err
	}

	transaction, err := t.trRepo.Create(ctx, transaction)
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (s TransactionService) CancelLatestOddRecords(ctx context.Context, limit int) error {
	transactions, err := s.trRepo.FindLatestOddRecords(ctx, limit)
	if err != nil {
		log.Printf("error find latest odd records %s", err)
		return err
	}

	for _, t := range transactions {
		if t.IsCanceled {
			continue
		}

		err := s.trRepo.MarkAsCanceled(ctx, t)
		if err != nil {
			log.Printf("error mark as canceled %s, %+v", err, t)
			return err
		}

		amount := t.Amount
		switch t.State {
		case domain.StateWin:
			amount = -amount
		case domain.StateLost:
			amount = +amount
		}

		err = s.acService.UpdateBalance(ctx, t.AccountID, amount)
		if err != nil {
			log.Printf("error mark as canceled %s, %+v", err, t)
			// return err
		}

		log.Println("account balance updated", t.AccountID, amount)
	}

	return nil
}

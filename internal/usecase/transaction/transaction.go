package transaction

import (
	"log"

	"github.com/asetriza/transaction-tracker/internal/domain"
	"github.com/asetriza/transaction-tracker/internal/usecase/account"
)

type TransactionRepository interface {
	FindLatestOddRecords(limit int) ([]domain.Transaction, error)
	MarkAsCanceled(transaction domain.Transaction) error
}

type TransactionService struct {
	trRepo    TransactionRepository
	acService account.Service
}

func NewTransactionService(
	trRepo TransactionRepository,
	acService account.Service,
) TransactionService {
	return TransactionService{
		trRepo:    trRepo,
		acService: acService,
	}
}

func (s TransactionService) CancelLatestOddRecords(limit int) error {
	transactions, err := s.trRepo.FindLatestOddRecords(limit)
	if err != nil {
		log.Printf("error find latest odd records %s", err)
		return err
	}

	for _, t := range transactions {
		if t.IsCanceled {
			continue
		}

		err := s.trRepo.MarkAsCanceled(t)
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

		err = s.acService.UpdateBalance(t.AccountID, amount)
		if err != nil {
			log.Printf("error mark as canceled %s, %+v", err, t)
			// return err
		}

		log.Println("account balance updated", t.AccountID, amount)
	}

	return nil
}

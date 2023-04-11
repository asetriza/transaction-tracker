package transaction

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

var _ TransactionRepository = (*MockTransactionRepository)(nil)

type (
	MockCreateFunc               func(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error)
	MockFindLatestOddRecordsFunc func(ctx context.Context, limit int) ([]domain.Transaction, error)
	MockMarkAsCanceledFunc       func(ctx context.Context, transaction domain.Transaction) error
)

type MockTransactionRepository struct {
	CreateFunc               MockCreateFunc
	FindLatestOddRecordsFunc MockFindLatestOddRecordsFunc
	MarkAsCanceledFunc       MockMarkAsCanceledFunc
}

func (r MockTransactionRepository) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	return r.CreateFunc(ctx, transaction)
}

func (r MockTransactionRepository) FindLatestOddRecords(ctx context.Context, limit int) ([]domain.Transaction, error) {
	return r.FindLatestOddRecordsFunc(ctx, limit)
}

func (r MockTransactionRepository) MarkAsCanceled(ctx context.Context, transaction domain.Transaction) error {
	return r.MarkAsCanceledFunc(ctx, transaction)
}

var _ AccountService = (*MockAccountService)(nil)

type (
	MockUpdateBalanceFunc func(ctx context.Context, id int, amount float64) error
	MockCreateAccountFunc func(ctx context.Context, account domain.Account) (domain.Account, error)
)

type MockAccountService struct {
	CreateFunc        MockCreateAccountFunc
	UpdateBalanceFunc MockUpdateBalanceFunc
}

func (r MockAccountService) UpdateBalance(ctx context.Context, id int, amount float64) error {
	return r.UpdateBalanceFunc(ctx, id, amount)
}

func (r MockAccountService) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	return r.CreateFunc(ctx, account)
}

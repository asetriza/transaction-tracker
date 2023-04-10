package postgresql

import (
	"context"
	"fmt"

	"github.com/asetriza/transaction-tracker/internal/tracker/domain"
	"github.com/lib/pq"
)

func (p PostgreSQL) CreateTransaction(ctx context.Context, tr domain.Transaction) (domain.Transaction, error) {
	query := `INSERT INTO transactions (state, amount, transaction_id) VALUES ($1, $2, $3)`

	_, err := p.db.ExecContext(ctx, query, tr.State, tr.Amount, tr.TransactionID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return domain.Transaction{}, fmt.Errorf("PostgreSQL error: %s", pgErr.Message)
		}
		return domain.Transaction{}, fmt.Errorf("database error: %s", err)
	}

	return tr, nil
}

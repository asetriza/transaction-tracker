package transactionrepo

import (
	"context"
	"errors"

	"github.com/asetriza/transaction-tracker/internal/domain"
	"github.com/asetriza/transaction-tracker/internal/repository/postgresql"
)

type transactionRepo struct {
	postgresql.Client
}

func New(client postgresql.Client) transactionRepo {
	return transactionRepo{client}
}

func (tr transactionRepo) Create(ctx context.Context, transaction domain.Transaction) (domain.Transaction, error) {
	tx, err := tr.DB.Begin()
	if err != nil {
		return domain.Transaction{}, err
	}

	stmt, err := tx.Prepare("INSERT INTO transactions (transaction_id, account_id, state, amount, created_at) VALUES ($1, $2, $3, $4, now())")
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		transaction.TransactionID,
		transaction.AccountID,
		transaction.State,
		transaction.Amount,
	)
	if err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	if rowsAffected != 1 {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	stmt, err = tx.Prepare("UPDATE accounts SET balance = balance + $1 WHERE id = $2")
	if err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}
	defer stmt.Close()

	switch transaction.State {
	case domain.StateWin:
		transaction.Amount = +transaction.Amount
	case domain.StateLost:
		transaction.Amount = -transaction.Amount
	}

	res, err = stmt.Exec(transaction.Amount, transaction.AccountID)
	if err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	if rowsAffected != 1 {
		tx.Rollback()
		return domain.Transaction{}, err
	}

	err = tx.Commit()
	if err != nil {
		return domain.Transaction{}, err
	}

	return transaction, nil
}

func (r transactionRepo) FindLatestOddRecords(ctx context.Context, limit int) ([]domain.Transaction, error) {
	query := `
		SELECT id, transaction_id, account_id, state, amount, created_at
		FROM transactions
		WHERE MOD(id, 2) = 1
		ORDER BY id DESC
		LIMIT $1
	`

	rows, err := r.DB.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]domain.Transaction, 0, limit)
	for rows.Next() {
		var t domain.Transaction
		err = rows.Scan(&t.ID, &t.TransactionID, &t.AccountID, &t.State, &t.Amount, &t.CreatedAt)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, errors.New("no odd records found")
	}

	return transactions, nil
}

func (tr transactionRepo) MarkAsCanceled(ctx context.Context, transaction domain.Transaction) error {
	query := `
		UPDATE transactions
		SET is_canceled = True
		WHERE id = $1
	`

	result, err := tr.DB.ExecContext(ctx, query, transaction.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("transaction already canceled or not found")
	}

	return nil
}

package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/asetriza/transaction-tracker/internal/domain"
)

func (p PostgreSQL) CreateTransaction(ctx context.Context, tr domain.Transaction) (domain.Transaction, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return domain.Transaction{}, err
	}

	stmt, err := tx.Prepare("INSERT INTO transactions (transaction_id, account_id, state, amount, created_at) VALUES ($1, $2, $3, $4, now())")
	if err != nil {
		return domain.Transaction{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(tr.TransactionID, tr.AccountID, tr.State, tr.Amount)
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

	res, err = stmt.Exec(tr.Amount, tr.AccountID)
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
		log.Fatal(err)
	}

	return tr, nil
}

func (r PostgreSQL) FindLatestOddRecords(limit int) ([]domain.Transaction, error) {
	query := `
		SELECT id, transaction_id, account_id, state, amount, created_at
		FROM transactions
		WHERE MOD(id, 2) = 1
		ORDER BY id DESC
		LIMIT $1
	`

	rows, err := r.db.Query(query, limit)
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

func (r PostgreSQL) MarkAsCanceled(transaction domain.Transaction) error {
	query := `
		UPDATE transactions
		SET is_canceled = True
		WHERE id = $1
	`

	result, err := r.db.Exec(query, transaction.ID)
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

func (r PostgreSQL) CreateAccount(ctx context.Context, acc domain.Account) (domain.Account, error) {
	account := domain.Account{}
	err := r.db.QueryRowContext(ctx, "INSERT INTO accounts (balance) VALUES ($1) RETURNING id, balance", acc.Balance).
		Scan(&acc.ID, &acc.Balance)
	if err != nil {
		log.Println("queryrow error", err.Error())
		return domain.Account{}, err
	}

	return account, nil
}

func (r PostgreSQL) FindById(id int) (domain.Account, error) {
	account := domain.Account{}

	row := r.db.QueryRow("SELECT id, balance FROM accounts WHERE id = $1", id)
	err := row.Scan(&account.ID, &account.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Account{}, domain.ErrAccountNotFound
		}

		return domain.Account{}, err
	}

	return account, nil
}

func (r PostgreSQL) UpdateBalance(account domain.Account) error {
	result, err := r.db.Exec("UPDATE accounts SET balance = $1 WHERE id = $2", account.Balance, account.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrAccountNotFound
	}

	return nil
}

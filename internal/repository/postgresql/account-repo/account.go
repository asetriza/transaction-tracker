package accountrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/asetriza/transaction-tracker/internal/domain"
	"github.com/asetriza/transaction-tracker/internal/repository/postgresql"
)

type accountRepo struct {
	postgresql.Client
}

func New(client postgresql.Client) accountRepo {
	return accountRepo{client}
}

func (ar accountRepo) Create(ctx context.Context, acc domain.Account) (domain.Account, error) {
	var id int
	var balance float64
	err := ar.DB.QueryRowContext(ctx, "INSERT INTO accounts (balance) VALUES ($1) RETURNING id, balance", acc.Balance).
		Scan(&id, &balance)
	if err != nil {
		log.Println("query row error", err.Error())
		return domain.Account{}, err
	}

	fmt.Println("new account", id, balance)

	return domain.Account{
		ID:      id,
		Balance: balance,
	}, nil
}

func (ar accountRepo) FindById(ctx context.Context, id int) (domain.Account, error) {
	account := domain.Account{}

	row := ar.DB.QueryRowContext(ctx, "SELECT id, balance FROM accounts WHERE id = $1", id)
	err := row.Scan(&account.ID, &account.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Account{}, domain.ErrAccountNotFound
		}

		return domain.Account{}, err
	}

	return account, nil
}

func (ar accountRepo) UpdateBalance(ctx context.Context, account domain.Account) error {
	result, err := ar.DB.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", account.Balance, account.ID)
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

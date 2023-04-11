package rest

import (
	"context"

	"github.com/asetriza/transaction-tracker/internal/domain"
	models "github.com/asetriza/transaction-tracker/internal/models"
)

func (h Handler) CreateAccount(ctx context.Context, req *models.CreateAccountReq) (models.CreateAccountRes, error) {
	dAccount := toDomainAccount(*req)

	transaction, err := h.account.CreateAccount(ctx, dAccount)
	if err != nil {
		return &models.CreateAccountBadRequest{}, err
	}

	return fromDomainAccount(transaction), nil
}

func toDomainAccount(ac models.CreateAccountReq) domain.Account {
	return domain.Account{
		Balance: ac.Balance,
	}
}

func fromDomainAccount(account domain.Account) *models.CreateAccountOK {
	return &models.CreateAccountOK{
		AccountId: account.ID,
		Balance:   account.Balance,
	}
}

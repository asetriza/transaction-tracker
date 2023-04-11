package rest

import (
	"context"
	"net/http"

	models "github.com/asetriza/transaction-tracker/internal/models"
	"github.com/asetriza/transaction-tracker/internal/usecase/account"
	"github.com/asetriza/transaction-tracker/internal/usecase/transaction"
)

type Handler struct {
	transaction transaction.Service
	account     account.Service
}

func NewHandler(
	transaction transaction.Service,
	account account.Service,
) Handler {
	return Handler{
		transaction: transaction,
		account:     account,
	}
}

func (h Handler) NewError(ctx context.Context, err error) *models.ErrorStatusCode {
	return &models.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: models.Error{
			Message: http.StatusText(http.StatusInternalServerError),
		},
	}
}

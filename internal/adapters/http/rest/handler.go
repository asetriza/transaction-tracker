package rest

import (
	"context"
	"net/http"

	models "github.com/asetriza/transaction-tracker/internal/models"
	"github.com/asetriza/transaction-tracker/internal/usecase/account"
	"github.com/asetriza/transaction-tracker/internal/usecase/tracker"
)

type Handler struct {
	tracker tracker.Service
	account account.Service
}

func NewHandler(
	tracker tracker.Service,
	account account.Service,
) Handler {
	return Handler{
		tracker: tracker,
		account: account,
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

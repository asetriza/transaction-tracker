package rest

import (
	"context"
	"net/http"

	models "github.com/asetriza/transaction-tracker/internal/tracker/models"
	"github.com/asetriza/transaction-tracker/internal/tracker/usecase/tracker"
)

type Handler struct {
	tracker tracker.Service
}

func NewHandler(tracker tracker.Service) Handler {
	return Handler{
		tracker: tracker,
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

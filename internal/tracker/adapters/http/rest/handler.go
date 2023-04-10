package rest

import "github.com/asetriza/transaction-tracker/internal/tracker/usecase/tracker"

type Handler struct {
	tracker tracker.Service
}

func NewHandler(tracker tracker.Service) Handler {
	return Handler{
		tracker: tracker,
	}
}

package domain

import "errors"

type Account struct {
	ID      int
	Balance float64
}

var ErrAccountNotFound = errors.New("account not found")

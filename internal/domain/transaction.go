package domain

import (
	"fmt"
	"time"
)

type State int

const (
	StateWin State = iota
	StateLost
)

var stateNames = map[State]string{
	StateWin:  "Win",
	StateLost: "Lost",
}

func (s State) String() string {
	return stateNames[s]
}

func StringToState(s string) (State, error) {
	for k, v := range stateNames {
		if v == s {
			return k, nil
		}
	}
	return 0, fmt.Errorf("invalid state: %s", s)
}

type TransactionID string

type Transaction struct {
	ID            int
	TransactionID TransactionID
	AccountID     int
	State         State
	Amount        float64
	IsCanceled    bool
	CreatedAt     time.Time
}

func (t Transaction) Validate() error {
	if t.Amount <= 0 {
		return NewValidationError("transaction amount should be greater than 0")
	}

	return nil
}

package domain

import "fmt"

type State int

const (
	Win State = iota
	Lost
)

var stateNames = map[State]string{
	Win:  "Win",
	Lost: "Lost",
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

type ID string

type Transaction struct {
	State         State
	Amount        float64
	TransactionID ID
}

func (t Transaction) Validate() error {
	if t.Amount <= 0 {
		return NewValidationError("transaction amount should be greater than 0")
	}

	return nil
}

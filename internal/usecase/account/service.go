package account

type Service interface {
	UpdateBalance(id int, amount float64) error
}

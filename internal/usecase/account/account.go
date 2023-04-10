package account

import "github.com/asetriza/transaction-tracker/internal/domain"

type AccountRepository interface {
	FindById(id int) (domain.Account, error)
	UpdateBalance(account domain.Account) error
}

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(repo AccountRepository) AccountService {
	return AccountService{repo}
}

func (s AccountService) UpdateBalance(id int, amount float64) error {
	account, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	account.Balance += amount

	return s.repo.UpdateBalance(account)
}

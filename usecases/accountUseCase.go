package usecases

import (
	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

type accountUseCase struct {
	accountRepository interfaces.AccountRepository
}

func NewAccountUseCase(accountRepository interfaces.AccountRepository) interfaces.AccountUseCase {
	return &accountUseCase{accountRepository}
}

func (a *accountUseCase) GetAccount(accountID int64) (*models.Account, error) {

	account, err := a.accountRepository.GetAccount(accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (a *accountUseCase) CreateNewAccount(account *models.Account) (*models.Account, error) {

	exists, err := a.accountRepository.ValidateAccount(account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, err
	}

	account.ID, err = a.accountRepository.CreateNewAccount(account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	return account, nil
}

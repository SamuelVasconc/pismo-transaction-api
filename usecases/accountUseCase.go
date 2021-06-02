package usecases

import (
	"errors"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

type accountUseCase struct {
	accountRepository interfaces.AccountRepository
}

//NewAccountUseCase ...
func NewAccountUseCase(accountRepository interfaces.AccountRepository) interfaces.AccountUseCase {
	return &accountUseCase{accountRepository}
}

//GetAccount by ID
func (a *accountUseCase) GetAccount(accountID int64) (*models.Account, error) {

	account, err := a.accountRepository.GetAccount(accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}

//CreateNewAccount by Document Number and generate an ID
func (a *accountUseCase) CreateNewAccount(account *models.Account) (*models.Account, error) {

	exists, err := a.accountRepository.ValidateAccount(account.DocumentNumber)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("This account already exists.")
	}

	if account.AvaliableCreditLimit < 0 {
		return nil, errors.New("The credit of this account is negative.")
	}

	account.ID, err = a.accountRepository.CreateNewAccount(account.DocumentNumber, account.AvaliableCreditLimit)
	if err != nil {
		return nil, err
	}

	return account, nil
}

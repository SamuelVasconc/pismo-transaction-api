package usecases

import (
	"errors"
	"time"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

type transactionUseCase struct {
	transactionRepository interfaces.TransactionRepository
	operationRepository   interfaces.OperationRepository
	accountRepository     interfaces.AccountRepository
}

//NewTransactionUseCase ...
func NewTransactionUseCase(transactionRepository interfaces.TransactionRepository, operationRepository interfaces.OperationRepository, accountRepository interfaces.AccountRepository) interfaces.TransactionUseCase {
	return &transactionUseCase{transactionRepository, operationRepository, accountRepository}
}

//CreateNewTransaction by request
func (t *transactionUseCase) CreateNewTransaction(transaction *models.Transaction) (*models.Transaction, error) {

	movementType, err := t.operationRepository.GetOperation(transaction.OperationTypeID)
	if err != nil {
		return nil, errors.New("This operation does not exists.")
	}

	_, err = t.accountRepository.GetAccount(transaction.AcountID)
	if err != nil {
		return nil, errors.New("This account does not exists. Please register the new account before any transaction.")
	}

	transaction.EventDate = time.Now()
	if movementType == "S" {
		transaction.Amount = transaction.Amount * (-1)
	}

	transaction.ID, err = t.transactionRepository.CreateNewTransaction(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

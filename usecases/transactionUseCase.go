package usecases

import (
	"time"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

type transactionUseCase struct {
	transactionRepository interfaces.TransactionRepository
	operationRepository   interfaces.OperationRepository
}

func NewTransactionUseCase(transactionRepository interfaces.TransactionRepository, operationRepository interfaces.OperationRepository) interfaces.TransactionUseCase {
	return &transactionUseCase{transactionRepository, operationRepository}
}

func (t *transactionUseCase) CreateNewTransaction(transaction *models.Transaction) (int64, error) {

	movementType, err := t.operationRepository.GetOperation(transaction.OperationTypeID)
	if err != nil {
		return 0, err
	}

	transaction.EventDate = time.Now()
	if movementType == "S" {
		transaction.Amount = transaction.Amount * (-1)
	}

	id, err := t.transactionRepository.CreateNewTransaction(transaction)
	if err != nil {
		return 0, err
	}

	return id, nil
}

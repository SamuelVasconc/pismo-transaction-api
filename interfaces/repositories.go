package interfaces

import (
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

//HealthCheckRepository ...
type HealthCheckRepository interface {
	Check() (*models.HealthCheck, error)
}

//AccountRepository ...
type AccountRepository interface {
	GetAccount(accountID int64) (*models.Account, error)
	CreateNewAccount(documentNumber string) (int64, error)
	ValidateAccount(documentNumber string) (bool, error)
}

//TransactionRepository ...
type TransactionRepository interface {
	CreateNewTransaction(transaction *models.Transaction) (int64, error)
}

//OperationRepository ...
type OperationRepository interface {
	GetOperation(id int64) (string, error)
}

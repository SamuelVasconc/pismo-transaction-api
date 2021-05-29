package interfaces

import (
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

//HealthCheckUseCase ...
type HealthCheckUseCase interface {
	Check() (*models.HealthCheck, error)
}

//AccountUseCase ...
type AccountUseCase interface {
	GetAccount(accountID int64) (*models.Account, error)
	CreateNewAccount(account *models.Account) (*models.Account, error)
}

//TransactionUseCase ...
type TransactionUseCase interface {
	CreateNewTransaction(transaction *models.Transaction) (*models.Transaction, error)
}

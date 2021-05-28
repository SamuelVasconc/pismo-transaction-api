package repositories

import (
	"database/sql"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

type transactionRepository struct {
	Conn *sql.DB
}

//NewTransactionRepository ...
func NewTransactionRepository(Conn *sql.DB) interfaces.TransactionRepository {
	return &transactionRepository{Conn}
}

func (a *transactionRepository) CreateNewTransaction(transaction *models.Transaction) (int64, error) {
	query := `INSERT INTO t_transactions (account_id, operation_type_id, amount, event_date) VALUES(?, ?, ?, ?)`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer smt.Close()

	result, err := smt.Exec(transaction.AcountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

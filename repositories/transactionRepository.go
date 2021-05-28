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
	query := `INSERT INTO t_transactions (transaction_id, account_id, operation_type_id, amount, event_date) VALUES(DEFAULT, $1, $2, $3, $4)
				RETURNING transaction_id`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer smt.Close()

	var id int64
	err = smt.QueryRow(transaction.AcountID, transaction.OperationTypeID, transaction.Amount, transaction.EventDate).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

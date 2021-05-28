package repositories_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/SamuelVasconc/pismo-transaction-api/models"

	"github.com/SamuelVasconc/pismo-transaction-api/repositories"

	"github.com/stretchr/testify/assert"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateNewTransaction(t *testing.T) {

	var transaction_id int64

	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := `INSERT INTO t_transactions (transaction_id, account_id, operation_type_id, amount, event_date) VALUES(DEFAULT, $1, $2, $3, $4)
	RETURNING transaction_id`

	rows := sqlmock.NewRows([]string{"transaction_id"}).AddRow(transaction_id)
	mockTransaction := &models.Transaction{
		AcountID:        1,
		Amount:          12.5,
		EventDate:       time.Now(),
		OperationTypeID: 3,
	}

	t.Run("success", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(mockTransaction.AcountID, mockTransaction.OperationTypeID, mockTransaction.Amount, mockTransaction.EventDate).WillReturnRows(rows)

		repository := repositories.NewTransactionRepository(db)

		_, err = repository.CreateNewTransaction(mockTransaction)
		assert.NoError(t, err)
	})

	t.Run("success no lines", func(t *testing.T) {
		errnorows := sql.ErrNoRows
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(mockTransaction.AcountID, mockTransaction.OperationTypeID, mockTransaction.Amount, mockTransaction.EventDate).WillReturnError(errnorows)

		repository := repositories.NewTransactionRepository(db)

		_, err = repository.CreateNewTransaction(mockTransaction)
		assert.Error(t, err)
		assert.Equal(t, err, errnorows)
	})

	t.Run("error-database", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectExec().WithArgs(mockTransaction.AcountID, mockTransaction.OperationTypeID, mockTransaction.Amount, mockTransaction.EventDate).WillReturnResult(sqlmock.NewResult(1, 1))
		mockSql.ExpectQuery(query).WithArgs().WillReturnRows(rows)
		db.Close()
		repository := repositories.NewTransactionRepository(db)

		_, err = repository.CreateNewTransaction(mockTransaction)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockSql.ExpectPrepare(query)
		mockSql.ExpectQuery(query).WillReturnError(errors.New("error"))

		repository := repositories.NewTransactionRepository(db)

		_, err = repository.CreateNewTransaction(mockTransaction)
		assert.Error(t, err)
	})
}

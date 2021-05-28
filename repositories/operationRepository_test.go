package repositories_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/SamuelVasconc/pismo-transaction-api/repositories"

	"github.com/stretchr/testify/assert"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetOperation(t *testing.T) {

	var id int64
	var movement_type string

	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := `SELECT movement_type FROM t_operation WHERE operation_type_id = $1`

	rows := sqlmock.NewRows([]string{"movement_type"}).AddRow(movement_type)
	id = 2

	t.Run("success", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(id).WillReturnRows(rows)

		repository := repositories.NewOperationRepository(db)

		_, err = repository.GetOperation(id)
		assert.NoError(t, err)
	})

	t.Run("success no lines", func(t *testing.T) {
		errnorows := sql.ErrNoRows
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(id).WillReturnError(errnorows)

		repository := repositories.NewOperationRepository(db)

		_, err = repository.GetOperation(id)
		assert.Error(t, err)
		assert.Equal(t, err, errnorows)
	})

	t.Run("error-database", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
		mockSql.ExpectQuery(query).WithArgs().WillReturnRows(rows)
		db.Close()
		repository := repositories.NewOperationRepository(db)

		_, err = repository.GetOperation(id)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockSql.ExpectPrepare(query)
		mockSql.ExpectQuery(query).WillReturnError(errors.New("error"))

		repository := repositories.NewOperationRepository(db)

		_, err = repository.GetOperation(id)
		assert.Error(t, err)
	})
}

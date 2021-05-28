package repositories_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/SamuelVasconc/pismo-transaction-api/repositories"

	"github.com/stretchr/testify/assert"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetAccount(t *testing.T) {

	var id int64
	var document_number string

	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := `SELECT document_number FROM t_accounts WHERE account_id = $1`

	rows := sqlmock.NewRows([]string{"document_number"}).AddRow(document_number)
	id = 1234

	t.Run("success", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(id).WillReturnRows(rows)

		repository := repositories.NewAccountRepository(db)

		_, err = repository.GetAccount(id)
		assert.NoError(t, err)
	})

	t.Run("success no lines", func(t *testing.T) {
		errnorows := sql.ErrNoRows
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(id).WillReturnError(errnorows)

		repository := repositories.NewAccountRepository(db)

		_, err = repository.GetAccount(id)
		assert.Error(t, err)
		assert.Equal(t, err, errnorows)
	})

	t.Run("error-database", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
		mockSql.ExpectQuery(query).WithArgs().WillReturnRows(rows)
		db.Close()
		repository := repositories.NewAccountRepository(db)

		_, err = repository.GetAccount(id)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockSql.ExpectPrepare(query)
		mockSql.ExpectQuery(query).WillReturnError(errors.New("error"))

		repository := repositories.NewAccountRepository(db)

		_, err = repository.GetAccount(id)
		assert.Error(t, err)
	})
}

func TestValidateAccount(t *testing.T) {

	var document_number string

	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := `SELECT 1 FROM t_accounts WHERE document_number = $1`

	rows := sqlmock.NewRows([]string{"column"}).AddRow(1)
	document_number = "43212134"

	t.Run("success", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(document_number).WillReturnRows(rows)

		repository := repositories.NewAccountRepository(db)

		_, err = repository.ValidateAccount(document_number)
		assert.NoError(t, err)
	})

	t.Run("success no lines", func(t *testing.T) {
		errnorows := sql.ErrNoRows
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(document_number).WillReturnError(errnorows)

		repository := repositories.NewAccountRepository(db)

		_, err = repository.ValidateAccount(document_number)
		assert.NoError(t, err)
	})

	t.Run("error-database", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectExec().WithArgs(document_number).WillReturnResult(sqlmock.NewResult(1, 1))
		mockSql.ExpectQuery(query).WithArgs().WillReturnRows(rows)
		db.Close()
		repository := repositories.NewAccountRepository(db)

		_, err = repository.ValidateAccount(document_number)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockSql.ExpectPrepare(query)
		mockSql.ExpectQuery(query).WillReturnError(errors.New("error"))

		repository := repositories.NewAccountRepository(db)

		_, err = repository.ValidateAccount(document_number)
		assert.Error(t, err)
	})
}

func TestCreateNewAccount(t *testing.T) {

	var document_number string

	db, mockSql, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	query := `INSERT INTO t_accounts (account_id, document_number) VALUES(DEFAULT, $1) RETURNING account_id`

	rows := sqlmock.NewRows([]string{"account_id"}).AddRow(1)
	document_number = "43212134"

	t.Run("success", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(document_number).WillReturnRows(rows)

		repository := repositories.NewAccountRepository(db)

		_, err = repository.CreateNewAccount(document_number)
		assert.NoError(t, err)
	})

	t.Run("success no lines", func(t *testing.T) {
		errnorows := sql.ErrNoRows
		mockSql.ExpectPrepare(query).ExpectQuery().WithArgs(document_number).WillReturnError(errnorows)

		repository := repositories.NewAccountRepository(db)

		_, err = repository.CreateNewAccount(document_number)
		assert.Error(t, err)
		assert.Equal(t, err, errnorows)
	})

	t.Run("error-database", func(t *testing.T) {
		mockSql.ExpectPrepare(query).ExpectExec().WithArgs(document_number).WillReturnResult(sqlmock.NewResult(1, 1))
		mockSql.ExpectQuery(query).WithArgs().WillReturnRows(rows)
		db.Close()
		repository := repositories.NewAccountRepository(db)

		_, err = repository.CreateNewAccount(document_number)
		assert.Error(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockSql.ExpectPrepare(query)
		mockSql.ExpectQuery(query).WillReturnError(errors.New("error"))

		repository := repositories.NewAccountRepository(db)

		_, err = repository.CreateNewAccount(document_number)
		assert.Error(t, err)
	})
}

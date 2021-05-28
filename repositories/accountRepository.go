package repositories

import (
	"database/sql"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
	"github.com/SamuelVasconc/pismo-transaction-api/models"
)

type accountRepository struct {
	Conn *sql.DB
}

//NewAccountRepository ...
func NewAccountRepository(Conn *sql.DB) interfaces.AccountRepository {
	return &accountRepository{Conn}
}

func (a *accountRepository) GetAccount(accountID int64) (*models.Account, error) {
	query := `SELECT document_number FROM t_accounts WHERE account_id = ?`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer smt.Close()

	var accountNumber sql.NullString
	err = smt.QueryRow(accountID).Scan(&accountNumber)
	if err != nil {
		return nil, err
	}

	newAccount := &models.Account{
		ID:             accountID,
		DocumentNumber: accountNumber.String,
	}
	return newAccount, nil
}

func (a *accountRepository) CreateNewAccount(documentNumber string) (int64, error) {
	query := `INSERT INTO t_accounts (document_number) VALUES(?)`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer smt.Close()

	result, err := smt.Exec(&documentNumber)
	if err != nil {
		return 0, err
	}

	id, _ := result.LastInsertId()

	return id, nil
}

func (a *accountRepository) ValidateAccount(documentNumber string) (bool, error) {
	query := `SELECT 1 FROM t_accounts WHERE document_number = ?`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return false, err
	}

	defer smt.Close()

	var count sql.NullInt64
	err = smt.QueryRow(documentNumber).Scan(&count)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}

	return true, nil
}

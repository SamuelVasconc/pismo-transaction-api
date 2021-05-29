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

//GetAccount by ID
func (a *accountRepository) GetAccount(accountID int64) (*models.Account, error) {
	query := `SELECT document_number FROM t_accounts WHERE account_id = $1`

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

//CreateNewAccount by Document Number and generate an ID
func (a *accountRepository) CreateNewAccount(documentNumber string) (int64, error) {
	query := `INSERT INTO t_accounts (account_id, document_number) VALUES(DEFAULT, $1) RETURNING account_id`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer smt.Close()

	var id int64
	err = smt.QueryRow(&documentNumber).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

//ValidateAccount validate if the account already exists by Document Number
func (a *accountRepository) ValidateAccount(documentNumber string) (bool, error) {
	query := `SELECT 1 FROM t_accounts WHERE document_number = $1`

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

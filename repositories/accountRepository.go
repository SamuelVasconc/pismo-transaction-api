package repositories

import (
	"database/sql"
	"errors"

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
	query := `SELECT document_number, avaliable_credit_limit FROM t_accounts WHERE account_id = $1`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer smt.Close()

	var accountNumber sql.NullString
	var creditLimit sql.NullFloat64
	err = smt.QueryRow(accountID).Scan(&accountNumber, &creditLimit)
	if err != nil {
		return nil, err
	}

	newAccount := &models.Account{
		ID:                   accountID,
		DocumentNumber:       accountNumber.String,
		AvaliableCreditLimit: creditLimit.Float64,
	}
	return newAccount, nil
}

//CreateNewAccount by Document Number and generate an ID
func (a *accountRepository) CreateNewAccount(documentNumber string, avaliableCreditLimit float64) (int64, error) {
	query := `INSERT INTO t_accounts (account_id, document_number, avaliable_credit_limit) VALUES(DEFAULT, $1, $2) RETURNING account_id`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer smt.Close()

	var id int64
	err = smt.QueryRow(&documentNumber, avaliableCreditLimit).Scan(&id)
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

func (a *accountRepository) UpdateAccount(accountID int64, amount float64, movementType string) error {
	query := `SELECT avaliable_credit_limit FROM t_accounts WHERE account_id = $1`

	queryUpdate := `UPDATE t_accounts SET avaliable_credit_limit = $1 WHERE account_id = $2`

	smt, err := a.Conn.Prepare(query)
	if err != nil {
		return err
	}

	defer smt.Close()

	var limit float64
	err = smt.QueryRow(accountID).Scan(&limit)
	if err != nil {
		return err
	}

	if movementType == "S" {
		if amount > limit {
			return errors.New("limite Indisponivel para concluir a transação.")
		} else {
			newAmmount := limit - amount

			smtt, err := a.Conn.Prepare(queryUpdate)
			if err != nil {
				return err
			}

			defer smtt.Close()

			_, err = smtt.Exec(newAmmount, accountID)
			if err != nil {
				return err
			}
		}
	} else {

		newAmmount := limit + amount

		smtt, err := a.Conn.Prepare(queryUpdate)
		if err != nil {
			return err
		}

		defer smtt.Close()

		_, err = smtt.Exec(newAmmount, accountID)
		if err != nil {
			return err
		}
	}

	return nil
}

package repositories

import (
	"database/sql"

	"github.com/SamuelVasconc/pismo-transaction-api/interfaces"
)

type operationRepository struct {
	Conn *sql.DB
}

//NewOperationRepository ...
func NewOperationRepository(Conn *sql.DB) interfaces.OperationRepository {
	return &operationRepository{Conn}
}

func (o *operationRepository) GetOperation(id int64) (string, error) {
	query := `SELECT movement_type FROM t_operation WHERE operation_type_id = $1`

	smt, err := o.Conn.Prepare(query)
	if err != nil {
		return "", err
	}

	defer smt.Close()

	var movementType sql.NullString
	err = smt.QueryRow(id).Scan(&movementType)
	if err != nil {
		return "", err
	}

	return movementType.String, nil
}

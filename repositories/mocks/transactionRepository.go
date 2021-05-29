package mocks

import (
	"github.com/SamuelVasconc/pismo-transaction-api/models"
	"github.com/stretchr/testify/mock"
)

//TransactionRepository ...
type TransactionRepository struct {
	mock.Mock
}

//CreateNewTransaction ...
func (t *TransactionRepository) CreateNewTransaction(transaction *models.Transaction) (int64, error) {
	ret := t.Called(transaction)
	var r0 int64
	if rf, ok := ret.Get(0).(func(*models.Transaction) int64); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

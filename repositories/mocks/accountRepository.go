package mocks

import (
	"github.com/SamuelVasconc/pismo-transaction-api/models"
	"github.com/stretchr/testify/mock"
)

//AccountRepository ...
type AccountRepository struct {
	mock.Mock
}

//GetAccount ...
func (a *AccountRepository) GetAccount(id int64) (*models.Account, error) {
	ret := a.Called()
	var r0 *models.Account
	if rf, ok := ret.Get(0).(func(int64) *models.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//CreateNewAccount ...
func (a *AccountRepository) CreateNewAccount(documentNumber string) (int64, error) {
	ret := a.Called()
	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(documentNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(documentNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//ValidateAccount ...
func (a *AccountRepository) ValidateAccount(documentNumber string) (bool, error) {
	ret := a.Called()
	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(documentNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(documentNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

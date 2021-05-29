package mocks

import (
	"github.com/stretchr/testify/mock"
)

//OperationRepository ...
type OperationRepository struct {
	mock.Mock
}

//GetOperation ...
func (o *OperationRepository) GetOperation(id int64) (string, error) {
	ret := o.Called()
	var r0 string
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(string)
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

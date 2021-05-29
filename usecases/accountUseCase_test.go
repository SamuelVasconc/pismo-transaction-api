package usecases_test

import (
	"errors"
	"testing"

	"github.com/SamuelVasconc/pismo-transaction-api/models"

	"github.com/SamuelVasconc/pismo-transaction-api/repositories/mocks"
	"github.com/SamuelVasconc/pismo-transaction-api/usecases"

	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	var id int64
	mockRepo := new(mocks.AccountRepository)

	t.Run("success", func(t *testing.T) {
		id = 123
		mockRepo.On("GetAccount").Return(nil, nil).Once()

		u := usecases.NewAccountUseCase(mockRepo)
		_, err := u.GetAccount(id)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		id = 123
		mockRepo.On("GetAccount").Return(nil, errors.New("mock error")).Once()

		u := usecases.NewAccountUseCase(mockRepo)
		_, err := u.GetAccount(id)
		assert.Error(t, err)
	})
}

func TestCreateNewAccount(t *testing.T) {
	account := &models.Account{
		ID:             123,
		DocumentNumber: "4321",
	}
	mockRepo := new(mocks.AccountRepository)

	t.Run("success", func(t *testing.T) {
		mockRepo.On("ValidateAccount").Return(nil, nil).Once()
		mockRepo.On("CreateNewAccount").Return(nil, nil).Once()

		u := usecases.NewAccountUseCase(mockRepo)
		_, err := u.CreateNewAccount(account)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("ValidateAccount").Return(nil, nil).Once()
		mockRepo.On("CreateNewAccount").Return(nil, errors.New("mock error")).Once()

		u := usecases.NewAccountUseCase(mockRepo)
		_, err := u.CreateNewAccount(account)
		assert.Error(t, err)
	})

	t.Run("validate exists", func(t *testing.T) {
		mockRepo.On("ValidateAccount").Return(true, nil).Once()
		mockRepo.On("CreateNewAccount").Return(nil, nil).Once()

		u := usecases.NewAccountUseCase(mockRepo)
		_, err := u.CreateNewAccount(account)
		assert.Error(t, err)
	})

	t.Run("error-validate", func(t *testing.T) {
		mockRepo.On("ValidateAccount").Return(nil, errors.New("mock error")).Once()
		mockRepo.On("CreateNewAccount").Return(nil, nil).Once()

		u := usecases.NewAccountUseCase(mockRepo)
		_, err := u.CreateNewAccount(account)
		assert.Error(t, err)
	})
}

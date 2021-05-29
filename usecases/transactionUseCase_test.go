package usecases_test

import (
	"errors"
	"testing"
	"time"

	"github.com/SamuelVasconc/pismo-transaction-api/models"

	"github.com/SamuelVasconc/pismo-transaction-api/repositories/mocks"
	"github.com/SamuelVasconc/pismo-transaction-api/usecases"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	mockTransactionRepository := new(mocks.TransactionRepository)
	mockOperationRepository := new(mocks.OperationRepository)
	mockAccountRepository := new(mocks.AccountRepository)

	mocktransaction := &models.Transaction{
		AcountID:        123,
		Amount:          15.2,
		EventDate:       time.Now(),
		OperationTypeID: 3,
	}
	t.Run("success", func(t *testing.T) {
		mockOperationRepository.On("GetOperation").Return(nil, nil).Once()
		mockAccountRepository.On("GetAccount").Return(nil, nil).Once()
		mockTransactionRepository.On("CreateNewTransaction", mocktransaction).Return(nil, nil).Once()

		u := usecases.NewTransactionUseCase(mockTransactionRepository, mockOperationRepository, mockAccountRepository)
		_, err := u.CreateNewTransaction(mocktransaction)
		assert.NoError(t, err)
	})

	t.Run("success out event", func(t *testing.T) {
		mockOperationRepository.On("GetOperation").Return("S", nil).Once()
		mockAccountRepository.On("GetAccount").Return(nil, nil).Once()
		mockTransactionRepository.On("CreateNewTransaction", mocktransaction).Return(nil, nil).Once()

		u := usecases.NewTransactionUseCase(mockTransactionRepository, mockOperationRepository, mockAccountRepository)
		_, err := u.CreateNewTransaction(mocktransaction)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockOperationRepository.On("GetOperation").Return(nil, nil).Once()
		mockAccountRepository.On("GetAccount").Return(nil, nil).Once()
		mockTransactionRepository.On("CreateNewTransaction", mocktransaction).Return(nil, errors.New("mock error")).Once()

		u := usecases.NewTransactionUseCase(mockTransactionRepository, mockOperationRepository, mockAccountRepository)
		_, err := u.CreateNewTransaction(mocktransaction)
		assert.Error(t, err)
	})

	t.Run("error-account", func(t *testing.T) {
		mockAccountRepository.On("GetAccount").Return(nil, errors.New("mock error")).Once()
		mockOperationRepository.On("GetOperation").Return(nil, nil).Once()
		mockTransactionRepository.On("CreateNewTransaction", mocktransaction).Return(nil, nil).Once()

		u := usecases.NewTransactionUseCase(mockTransactionRepository, mockOperationRepository, mockAccountRepository)
		_, err := u.CreateNewTransaction(mocktransaction)
		assert.Error(t, err)
	})

	t.Run("error-operation", func(t *testing.T) {
		mockOperationRepository.On("GetOperation").Return(nil, errors.New("mock error")).Once()
		mockAccountRepository.On("GetAccount").Return(nil, nil).Once()
		mockTransactionRepository.On("CreateNewTransaction", mocktransaction).Return(nil, nil).Once()

		u := usecases.NewTransactionUseCase(mockTransactionRepository, mockOperationRepository, mockAccountRepository)
		_, err := u.CreateNewTransaction(mocktransaction)
		assert.Error(t, err)
	})
}

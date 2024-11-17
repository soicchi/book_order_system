package shippingAddresses

import (
	"fmt"
	"testing"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateShippingAddress(t *testing.T) {
	customerID, _ := uuid.NewV7()
	hashedPassword, _ := values.NewPassword("password")
	now := time.Now()
	customer := entity.ReconstructCustomer(
		customerID,
		"test",
		"test@test.co.jp",
		hashedPassword,
		&now,
		&now,
	)

	tests := []struct {
		name     string
		input    *dto.CreateShippingAddressInput
		mockFunc func(*interfaces.MockShippingAddressRepository, *interfaces.MockCustomerRepository)
		wantErr  bool
	}{
		{
			name: "create shipping address successfully",
			input: &dto.CreateShippingAddressInput{
				Prefecture: "Tokyo",
				City:       "Shinjuku",
				State:      "Nishishinjuku",
				CustomerID: customerID.String(),
			},
			mockFunc: func(shippingMock *interfaces.MockShippingAddressRepository, customerMock *interfaces.MockCustomerRepository) {
				customerMock.On("FetchByID", mock.Anything, mock.Anything).Return(customer, nil)
				shippingMock.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to create shipping address entity",
			input: &dto.CreateShippingAddressInput{
				Prefecture: "Tokyo",
				City:       "Shinjuku",
				State:      "Nishishinjuku",
				CustomerID: "invalid-customer-id",
			},
			mockFunc: func(shippingMock *interfaces.MockShippingAddressRepository, customerMock *interfaces.MockCustomerRepository) {
			},
			wantErr: true,
		},
		{
			name: "failed to fetch customer",
			input: &dto.CreateShippingAddressInput{
				Prefecture: "Tokyo",
				City:       "Shinjuku",
				State:      "Nishishinjuku",
				CustomerID: customerID.String(),
			},
			mockFunc: func(shippingMock *interfaces.MockShippingAddressRepository, customerMock *interfaces.MockCustomerRepository) {
				customerMock.On("FetchByID", mock.Anything, mock.Anything).Return(&entity.Customer{}, errors.NewCustomError(
					fmt.Errorf("failed to fetch customer"),
					errors.InternalServerError,
				))
			},
			wantErr: true,
		},
		{
			name: "failed to create shipping address",
			input: &dto.CreateShippingAddressInput{
				Prefecture: "Tokyo",
				City:       "Shinjuku",
				State:      "Nishishinjuku",
				CustomerID: customerID.String(),
			},
			mockFunc: func(shippingMock *interfaces.MockShippingAddressRepository, customerMock *interfaces.MockCustomerRepository) {
				customerMock.On("FetchByID", mock.Anything, mock.Anything).Return(customer, nil)
				shippingMock.On("Create", mock.Anything, mock.Anything).Return(errors.NewCustomError(
					fmt.Errorf("failed to create shipping address"),
					errors.InternalServerError,
				))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockShippingAddressRepo := interfaces.NewMockShippingAddressRepository()
			mockCustomerRepo := interfaces.NewMockCustomerRepository()
			tt.mockFunc(
				mockShippingAddressRepo,
				mockCustomerRepo,
			)

			e := echo.New()
			ctx := e.NewContext(nil, nil)

			logger := logging.NewMockLogger()
			useCase := NewShippingAddressUseCase(mockShippingAddressRepo, mockCustomerRepo, logger)
			err := useCase.CreateShippingAddress(ctx, tt.input)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

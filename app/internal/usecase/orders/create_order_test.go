package orders

import (
	"fmt"
	"testing"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/usecase/dto"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	customerID, _ := uuid.NewV7()
	shippingAddressID, _ := uuid.NewV7()
	now := time.Now()
	customerEntity := entity.ReconstructCustomer(customerID, "test", "test@test.co.jp", "hashed_password", &now, &now)
	shippingAddressEntity := entity.ReconstructShippingAddress(shippingAddressID, "tokyo", "shinjuku", "1-1", now, now, customerID)

	tests := []struct {
		name     string
		dto      *dto.OrderInput
		mockFunc func(
			*interfaces.MockOrderRepository,
			*interfaces.MockCustomerRepository,
			*interfaces.MockShippingAddressRepository,
		)
		wantErr bool
	}{
		{
			name: "create order successfully",
			dto: &dto.OrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *interfaces.MockOrderRepository,
				customerRepo *interfaces.MockCustomerRepository,
				shippingAddressRepo *interfaces.MockShippingAddressRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(shippingAddressEntity, nil)
				orderRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to generate customer UUID",
			dto: &dto.OrderInput{
				CustomerID:        "invalid_uuid",
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *interfaces.MockOrderRepository,
				customerRepo *interfaces.MockCustomerRepository,
				shippingAddressRepo *interfaces.MockShippingAddressRepository,
			) {
			},
			wantErr: true,
		},
		{
			name: "failed to fetch customer by ID",
			dto: &dto.OrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *interfaces.MockOrderRepository,
				customerRepo *interfaces.MockCustomerRepository,
				shippingAddressRepo *interfaces.MockShippingAddressRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, errors.NewCustomError(
					fmt.Errorf("failed to fetch customer by ID"),
					errors.InternalServerError,
				))
			},
			wantErr: true,
		},
		{
			name: "customer not found",
			dto: &dto.OrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *interfaces.MockOrderRepository,
				customerRepo *interfaces.MockCustomerRepository,
				shippingAddressRepo *interfaces.MockShippingAddressRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "failed to fetch shipping address by ID",
			dto: &dto.OrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *interfaces.MockOrderRepository,
				customerRepo *interfaces.MockCustomerRepository,
				shippingAddressRepo *interfaces.MockShippingAddressRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, errors.NewCustomError(
					fmt.Errorf("failed to fetch shipping address by ID"),
					errors.InternalServerError,
				))
			},
			wantErr: true,
		},
		{
			name: "shipping address not found",
			dto: &dto.OrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *interfaces.MockOrderRepository,
				customerRepo *interfaces.MockCustomerRepository,
				shippingAddressRepo *interfaces.MockShippingAddressRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "failed to create order",
			dto: &dto.OrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *interfaces.MockOrderRepository,
				customerRepo *interfaces.MockCustomerRepository,
				shippingAddressRepo *interfaces.MockShippingAddressRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(shippingAddressEntity, nil)
				orderRepo.On("Create", mock.Anything, mock.Anything).Return(errors.NewCustomError(
					fmt.Errorf("failed to create order"),
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

			// setup mock
			logger := logging.NewMockLogger()
			customerRepo := interfaces.NewMockCustomerRepository()
			shippingAddressRepo := interfaces.NewMockShippingAddressRepository()
			orderRepo := interfaces.NewMockOrderRepository()
			tt.mockFunc(orderRepo, customerRepo, shippingAddressRepo)

			// setup context
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			useCase := NewOrderUseCase(orderRepo, customerRepo, shippingAddressRepo, logger)
			err := useCase.CreateOrder(ctx, tt.dto)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

package orders

import (
	"fmt"
	"testing"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/shippingAddress"
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
	customerEntity := customer.Reconstruct(customerID, "test", "test@test.co.jp", "hashed_password", now, now)
	shippingAddressEntity := shippingAddress.Reconstruct(shippingAddressID, "tokyo", "shinjuku", "1-1", now, now, customerID)

	tests := []struct {
		name     string
		dto      *dto.CreateOrderInput
		mockFunc func(
			*order.MockRepository,
			*customer.MockRepository,
			*shippingAddress.MockRepository,
		)
		wantErr bool
	}{
		{
			name: "create order successfully",
			dto: &dto.CreateOrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(shippingAddressEntity, nil)
				orderRepo.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to generate customer UUID",
			dto: &dto.CreateOrderInput{
				CustomerID:        "invalid_uuid",
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
			) {
			},
			wantErr: true,
		},
		{
			name: "failed to fetch customer by ID",
			dto: &dto.CreateOrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
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
			dto: &dto.CreateOrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "failed to fetch shipping address by ID",
			dto: &dto.CreateOrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
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
			dto: &dto.CreateOrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)
			},
			wantErr: true,
		},
		{
			name: "failed to create order",
			dto: &dto.CreateOrderInput{
				CustomerID:        customerID.String(),
				ShippingAddressID: shippingAddressID.String(),
			},
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
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

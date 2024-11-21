package orders

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/domain/order"
	"github.com/soicchi/book_order_system/internal/domain/shippingAddress"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/presentation/validator"
	"github.com/soicchi/book_order_system/internal/usecase/orders"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	customerID, _ := uuid.NewV7()
	ShippingAddressID, _ := uuid.NewV7()
	now := time.Now()
	customerEntity := customer.Reconstruct(
		customerID,
		"test",
		"test@test.co.jp",
		"hashed_password",
		&now,
		&now,
	)
	shippingAddressEntity := shippingAddress.Reconstruct(
		ShippingAddressID,
		"tokyo",
		"shinjuku",
		"1-1",
		&now,
		&now,
	)

	tests := []struct {
		name        string
		requestBody string
		mockFunc    func(
			*order.MockRepository,
			*customer.MockRepository,
			*shippingAddress.MockRepository,
			*logging.MockLogger,
		)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:        "create order successfully",
			requestBody: `{"shipping_address_id": "` + ShippingAddressID.String() + `"}`,
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
				ml *logging.MockLogger,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(shippingAddressEntity, nil)
				orderRepo.On("Create", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus:   http.StatusCreated,
			expectedResponse: `{"message":"created order successfully"}`,
		},
		{
			name: "failed to create order with not found customer",
			requestBody: `{
				"shipping_address_id": "` + ShippingAddressID.String() + `"
			}`,
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
				ml *logging.MockLogger,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusNotFound,
			expectedResponse: `{
				"code": "not_found",
				"message": "Resource not found",
				"details": {
					"customer_id": "not found"
				}
			}`,
		},
		{
			name: "failed to create order with not found shipping address",
			requestBody: `{
				"shipping_address_id": "` + ShippingAddressID.String() + `"
			}`,
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
				ml *logging.MockLogger,
			) {
				customerRepo.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
				shippingAddressRepo.On("FetchByID", mock.Anything, mock.Anything).Return(nil, nil)
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusNotFound,
			expectedResponse: `{
				"code": "not_found",
				"message": "Resource not found",
				"details": {
					"shipping_address_id": "not found"
				}
			}`,
		},
		{
			name: "failed to create order with validation error",
			requestBody: `{
				"shipping_address_id": ""
			}`,
			mockFunc: func(
				orderRepo *order.MockRepository,
				customerRepo *customer.MockRepository,
				shippingAddressRepo *shippingAddress.MockRepository,
				ml *logging.MockLogger,
			) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"message": "Invalid request parameters",
				"details": {
					"ShippingAddressID": "required"
				}
			}`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			// setup request
			e := echo.New()
			e.Validator = validator.NewCustomValidator()
			req := httptest.NewRequest(http.MethodPost, "/customers/"+customerID.String()+"/orders", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("customer_id")
			ctx.SetParamValues(customerID.String())

			// setup mock
			logger := logging.NewMockLogger()
			orderRepo := order.NewMockRepository()
			customerRepo := customer.NewMockRepository()
			shippingAddressRepo := shippingAddress.NewMockRepository()
			tt.mockFunc(orderRepo, customerRepo, shippingAddressRepo, logger)

			useCase := orders.NewOrderUseCase(orderRepo, customerRepo, shippingAddressRepo, logger)
			handler := NewOrderHandler(useCase, logger)
			err := handler.CreateOrder(ctx)
			t.Log(err)

			res := rec.Result()
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			assert.JSONEq(t, tt.expectedResponse, string(body))
		})
	}
}

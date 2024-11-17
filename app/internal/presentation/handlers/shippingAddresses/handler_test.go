package shippingAddresses

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/presentation/validator"
	"github.com/soicchi/book_order_system/internal/usecase/shippingAddresses"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateShippingAddress(t *testing.T) {
	customerID, _ := uuid.NewV7()

	tests := []struct {
		name              string
		requestCustomerID string
		requestBody       string
		mockFunc          func(
			*interfaces.MockShippingAddressRepository,
			*interfaces.MockCustomerRepository,
			*logging.MockLogger,
		)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:              "create shipping address successfully",
			requestCustomerID: customerID.String(),
			requestBody: `{
				"prefecture": "Tokyo",
				"city": "Shinjuku",
				"state": "Nishishinjuku"
			}`,
			mockFunc: func(
				shippingMock *interfaces.MockShippingAddressRepository,
				customerMock *interfaces.MockCustomerRepository,
				ml *logging.MockLogger,
			) {
				customerMock.On("FetchByID", mock.Anything, mock.Anything).Return(&entity.Customer{}, nil)
				shippingMock.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedResponse: `{
				"message": "created shipping address successfully"
			}`,
		},
		{
			name:              "failed to create shipping address with empty prefecture",
			requestCustomerID: customerID.String(),
			requestBody: `{
				"prefecture": "",
				"city": "Shinjuku",
				"state": "Nishishinjuku"
			}`,
			mockFunc: func(
				shippingMock *interfaces.MockShippingAddressRepository,
				customerMock *interfaces.MockCustomerRepository,
				ml *logging.MockLogger,
			) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"prefecture": "required"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name:              "failed to create shipping address with empty city",
			requestCustomerID: customerID.String(),
			requestBody: `{
				"prefecture": "Tokyo",
				"city": "",
				"state": "Nishishinjuku"
			}`,
			mockFunc: func(
				shippingMock *interfaces.MockShippingAddressRepository,
				customerMock *interfaces.MockCustomerRepository,
				ml *logging.MockLogger,
			) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"city": "required"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name:              "failed to create shipping address with empty state",
			requestCustomerID: customerID.String(),
			requestBody: `{
				"prefecture": "Tokyo",
				"city": "Shinjuku",
				"state": ""
			}`,
			mockFunc: func(
				shippingMock *interfaces.MockShippingAddressRepository,
				customerMock *interfaces.MockCustomerRepository,
				ml *logging.MockLogger,
			) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"state": "required"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name:              "failed to create shipping address with not found customer id",
			requestCustomerID: customerID.String(),
			requestBody: `{
				"prefecture": "Tokyo",
				"city": "Shinjuku",
				"state": "Nishishinjuku"
			}`,
			mockFunc: func(
				shippingMock *interfaces.MockShippingAddressRepository,
				customerMock *interfaces.MockCustomerRepository,
				ml *logging.MockLogger,
			) {
				customerMock.On("FetchByID", mock.Anything, mock.Anything).Return(&entity.Customer{}, errors.NewCustomError(
					fmt.Errorf("failed to fetch customer"),
					errors.NotFound,
					errors.WithNotFoundDetails("customer_id"),
				))
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusNotFound,
			expectedResponse: `{
				"code": "not_found",
				"details": {
					"customer_id": "not found"
				},
				"message": "Resource not found"
			}`,
		},
		{
			name:              "failed to create shipping address with internal server error",
			requestCustomerID: customerID.String(),
			requestBody: `{
				"prefecture": "Tokyo",
				"city": "Shinjuku",
				"state": "Nishishinjuku"
			}`,
			mockFunc: func(
				shippingMock *interfaces.MockShippingAddressRepository,
				customerMock *interfaces.MockCustomerRepository,
				ml *logging.MockLogger,
			) {
				customerMock.On("FetchByID", mock.Anything, mock.Anything).Return(&entity.Customer{}, nil)
				shippingMock.On("Create", mock.Anything, mock.Anything).Return(errors.NewCustomError(
					fmt.Errorf("failed to create shipping address"),
					errors.InternalServerError,
				))
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedResponse: `{
				"code": "internal_error",
				"message": "An internal server error occurred"
			}`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// setup request
			req := httptest.NewRequest(
				http.MethodPost,
				"/customers/"+tt.requestCustomerID+"/shipping_addresses",
				strings.NewReader(tt.requestBody),
			)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			e := echo.New()
			e.Validator = validator.NewCustomValidator()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("customer_id")
			ctx.SetParamValues(tt.requestCustomerID)

			// setup mock
			shippingMock := interfaces.NewMockShippingAddressRepository()
			customerMock := interfaces.NewMockCustomerRepository()
			logger := logging.NewMockLogger()
			tt.mockFunc(shippingMock, customerMock, logger)

			// setup handler
			useCase := shippingAddresses.NewShippingAddressUseCase(shippingMock, customerMock, logger)
			handler := NewShippingAddressHandler(useCase, logger)

			handler.CreateShippingAddress(ctx)
			res := rec.Result()
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			assert.JSONEq(t, tt.expectedResponse, string(body))
		})
	}
}

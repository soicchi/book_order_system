package customers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/customer"
	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/presentation/validator"
	"github.com/soicchi/book_order_system/internal/usecase/customers"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	tests := []struct {
		name             string
		requestBody      string
		mockFunc         func(*customer.MockRepository, *logging.MockLogger)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "request to create customer successfully",
			requestBody: `{
				"name": "test",
				"email": "test@test.co.jp",
				"password": "password"
			}`,
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				m.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusCreated,
			expectedResponse: `{
				"message": "created customer successfully"
			}`,
		},
		{
			name: "request to create customer with empty name",
			requestBody: `{
				"name": "",
				"email": "test@test.co.jp",
				"password": "password"
			}`,
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"Name": "required"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name: "request to create customer with empty email",
			requestBody: `{
				"name": "test",
				"email": "",
				"password": "password"
			}`,
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"Email": "required"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name: "request to create customer with empty password",
			requestBody: `{
				"name": "test",
				"email": "test@test.co.jp",
				"password": ""
			}`,
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"Password": "required"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name: "request to create customer with invalid email",
			requestBody: `{
				"name": "test",
				"email": "invalid",
				"password": "password"
			}`,
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"Email": "email"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name: "failed to create customer",
			requestBody: `{
				"name": "test",
				"email": "test@test.co.jp",
				"password": "password"
			}`,
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				m.On("Create", mock.Anything, mock.Anything).Return(errors.NewCustomError(
					fmt.Errorf("failed to create customer"),
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

			// make mock
			logger := logging.NewMockLogger()
			mockRepo := customer.NewMockRepository()
			tt.mockFunc(mockRepo, logger)

			useCase := customers.NewCustomerUseCase(mockRepo, logging.NewMockLogger())

			// setup request
			e := echo.New()
			e.Validator = validator.NewCustomValidator()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/customers", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			handler := NewCustomerHandler(useCase, logger)

			handler.CreateCustomer(ctx)

			res := rec.Result()
			body, _ := io.ReadAll(res.Body)

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
			assert.JSONEq(t, tt.expectedResponse, string(body))
		})
	}
}

func TestFetchCustomer(t *testing.T) {
	customerID, _ := uuid.NewV7()
	now := time.Now()
	hashedPassword, _ := values.NewPassword("password")
	customer := entity.ReconstructCustomer(
		customerID,
		"test",
		"test@test.co.jp",
		hashedPassword,
		now,
		now,
	)

	tests := []struct {
		name             string
		id               string
		mockFunc         func(*customer.MockRepository, *logging.MockLogger)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "fetch customer successfully",
			id:   customerID.String(),
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				m.On("FetchByID", mock.Anything, mock.Anything).Return(customer, nil)
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedResponse: `{
				"message": "fetched customer successfully",
				"customer": {
					"id": "` + customerID.String() + `",
					"name": "` + customer.Name() + `",
					"email": "` + customer.Email() + `",
					"created_at": "` + customer.CreatedAt().Format("2006-01-02 15:04:05") + `",
					"updated_at": "` + customer.UpdatedAt().Format("2006-01-02 15:04:05") + `"
				}
			}`,
		},
		{
			name: "fetch customer with empty id",
			id:   "",
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: `{
				"code": "invalid_request",
				"details": {
					"ID": "required"
				},
				"message": "Invalid request parameters"
			}`,
		},
		{
			name: "failed to fetch customer",
			id:   customerID.String(),
			mockFunc: func(m *customer.MockRepository, ml *logging.MockLogger) {
				m.On("FetchByID", mock.Anything, mock.Anything).Return(nil, errors.NewCustomError(
					fmt.Errorf("failed to fetch customer"),
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

			// make mock
			logger := logging.NewMockLogger()
			mockRepo := customer.NewMockRepository()
			tt.mockFunc(mockRepo, logger)

			useCase := customers.NewCustomerUseCase(mockRepo, logger)

			// setup request
			e := echo.New()
			e.Validator = validator.NewCustomValidator()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/customers/"+tt.id, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			ctx.SetParamNames("id")
			ctx.SetParamValues(tt.id)

			handler := NewCustomerHandler(useCase, logger)
			handler.FetchCustomer(ctx)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			assert.JSONEq(t, tt.expectedResponse, rec.Body.String())
		})
	}
}

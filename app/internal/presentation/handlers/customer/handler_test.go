package customer

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/logging"
	"github.com/soicchi/book_order_system/internal/presentation/validator"
	"github.com/soicchi/book_order_system/internal/usecase/customers"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	tests := []struct {
		name             string
		requestBody      string
		mockFunc         func(m *interfaces.MockCustomerRepository, ml *logging.MockLogger)
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
			mockFunc: func(m *interfaces.MockCustomerRepository, ml *logging.MockLogger) {
				m.On("Create", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus:   http.StatusCreated,
			expectedResponse: `"created customer successfully"`,
		},
		{
			name: "request to create customer with empty name",
			requestBody: `{
				"name": "",
				"email": "test@test.co.jp",
				"password": "password"
			}`,
			mockFunc: func(m *interfaces.MockCustomerRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `"validation error"`,
		},
		{
			name: "request to create customer with empty email",
			requestBody: `{
				"name": "test",
				"email": "",
				"password": "password"
			}`,
			mockFunc: func(m *interfaces.MockCustomerRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `"validation error"`,
		},
		{
			name: "request to create customer with empty password",
			requestBody: `{
				"name": "test",
				"email": "test@test.co.jp",
				"password": ""
			}`,
			mockFunc: func(m *interfaces.MockCustomerRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `"validation error"`,
		},
		{
			name: "request to create customer with invalid email",
			requestBody: `{
				"name": "test",
				"email": "invalid",
				"password": "password"
			}`,
			mockFunc: func(m *interfaces.MockCustomerRepository, ml *logging.MockLogger) {
				ml.On("Error", mock.Anything, mock.Anything).Return(nil)
			},
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: `"validation error"`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// make mock
			logger := logging.NewMockLogger()
			mockRepo := interfaces.NewMockCustomerRepository()
			tt.mockFunc(mockRepo, logger)

			useCase := customers.NewCustomerUseCase(mockRepo, logging.NewMockLogger())

			e := echo.New()
			e.Validator = validator.NewCustomValidator()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/customers", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			handler := NewCustomerHandler(useCase, logger)

			if assert.NoError(t, handler.CreateCustomer(ctx)) {
				assert.Equal(t, tt.expectedStatus, rec.Code)
				assert.JSONEq(t, tt.expectedResponse, rec.Body.String())
			}
		})
	}
}

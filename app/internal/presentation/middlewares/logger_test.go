package middlewares

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"event_system/internal/logging"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCustomBodyDump(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		mockFunc func(*testing.T, *logging.MockLogger)
	}{
		{
			name: "test custom body dump",
			body: `{
				"name": "test",
				"email": "test@test.co.jp",
				"password": "password"
			}`,
			mockFunc: func(t *testing.T, m *logging.MockLogger) {
				m.On("LogAttrs", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			logger := logging.NewMockLogger()
			tt.mockFunc(t, logger)

			e := echo.New()

			// set middleware
			e.Use(middleware.BodyDump(CustomBodyDump(logger)))

			// request handler for test
			e.POST("/test", func(c echo.Context) error {
				return c.JSON(200, map[string]interface{}{"message": "ok"})
			})

			// set up test request
			req := httptest.NewRequest(http.MethodPost, "/test", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)

			logger.AssertExpectations(t)
		})
	}
}

func TestShouldMaskField(t *testing.T) {
	tests := []struct {
		name   string
		field  string
		expect bool
	}{
		{
			name:   "password",
			field:  "password",
			expect: true,
		},
		{
			name:   "email",
			field:  "email",
			expect: true,
		},
		{
			name:   "name",
			field:  "name",
			expect: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expect, shouldMaskField(tt.field))
		})
	}
}

func TestMaskField(t *testing.T) {
	tests := []struct {
		name           string
		fields         map[string]interface{}
		expectedFields map[string]interface{}
	}{
		{
			name: "mask email and password",
			fields: map[string]interface{}{
				"name":     "test",
				"email":    "test@test.co.jp",
				"password": "password",
			},
			expectedFields: map[string]interface{}{
				"name":     "test",
				"email":    "*****",
				"password": "*****",
			},
		},
		{
			name: "mask nested fields",
			fields: map[string]interface{}{
				"customer": map[string]interface{}{
					"name":     "test",
					"email":    "test@test.co.jp",
					"password": "password",
				},
			},
			expectedFields: map[string]interface{}{
				"customer": map[string]interface{}{
					"name":     "test",
					"email":    "*****",
					"password": "*****",
				},
			},
		},
		{
			name:           "body is nil",
			fields:         map[string]interface{}{},
			expectedFields: map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expectedFields, maskFields(tt.fields))
		})
	}
}

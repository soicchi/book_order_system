package customers

import (
	"errors"
	"testing"

	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/dto"
	"github.com/soicchi/book_order_system/internal/logger"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	tests := []struct {
		name     string
		dto      dto.CreateCustomerInput
		mockFunc func(*interfaces.MockCustomerRepository)
		wantErr  bool
	}{
		{
			name: "create customer successfully",
			dto: dto.CreateCustomerInput{
				Name:     "test",
				Email:    "test@test.co.jp",
				Password: "password",
			},
			mockFunc: func(m *interfaces.MockCustomerRepository) {
				m.On("Create", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to create customer entity with invalid password",
			dto: dto.CreateCustomerInput{
				Name:     "test",
				Email:    "test@test.co.jp",
				Password: "pass",
			},
			mockFunc: func(m *interfaces.MockCustomerRepository) {},
			wantErr:  true,
		},
		{
			name: "failed to create customer",
			dto: dto.CreateCustomerInput{
				Name:     "test",
				Email:    "test@test.co.jp",
				Password: "password",
			},
			mockFunc: func(m *interfaces.MockCustomerRepository) {
				m.On("Create", mock.Anything).Return(errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			logger := logger.NewMockLogger()
			repo := interfaces.NewMockCustomerRepository()
			useCase := NewCustomerUseCase(repo, logger)
			tt.mockFunc(repo)

			ctx := echo.New().NewContext(nil, nil)

			err := useCase.Execute(ctx, tt.dto)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

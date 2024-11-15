package customers

import (
	"errors"
	"testing"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/entity"
	"github.com/soicchi/book_order_system/internal/domain/interfaces"
	"github.com/soicchi/book_order_system/internal/domain/values"
	"github.com/soicchi/book_order_system/internal/logging"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchCustomer(t *testing.T) {
	customerID, _ := uuid.NewV7()
	hashedPassword, _ := values.NewPassword("password")
	now := time.Now()
	customerEntity := entity.ReconstructCustomer(
		customerID,
		"test",
		"test@test.com",
		hashedPassword,
		&now,
		&now,
	)

	tests := []struct {
		name     string
		id       string
		mockFunc func(*testing.T, *interfaces.MockCustomerRepository)
		wantErr  bool
	}{
		{
			name: "fetch customer successfully",
			id:   customerID.String(),
			mockFunc: func(t *testing.T, m *interfaces.MockCustomerRepository) {
				m.On("FetchByID", mock.Anything, mock.Anything).Return(customerEntity, nil)
			},
			wantErr: false,
		},
		{
			name: "failed to fetch customer",
			id:   customerID.String(),
			mockFunc: func(t *testing.T, m *interfaces.MockCustomerRepository) {
				m.On("FetchByID", mock.Anything, mock.Anything).Return(&entity.Customer{}, errors.New("error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Setup Mock
			repo := interfaces.NewMockCustomerRepository()
			tt.mockFunc(t, repo)

			// Setup context
			ctx := echo.New().NewContext(nil, nil)

			usecase := NewFetchCustomerUseCase(repo, logging.NewMockLogger())
			dto, err := usecase.Execute(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, dto.ID, customerID.String())
				assert.Equal(t, dto.Name, customerEntity.Name())
				assert.Equal(t, dto.Email, customerEntity.Email())
				assert.Equal(t, dto.CreatedAt, now)
				assert.Equal(t, dto.UpdatedAt, now)
			}
		})
	}
}

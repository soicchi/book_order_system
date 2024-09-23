package customer

import (
	"errors"
	"testing"

	customerDomain "github.com/soicchi/book_order_system/domain/customer"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomer(t *testing.T) {
	mockRepo := customerDomain.NewMockRepository()
	useCase := NewCreateCustomerUseCase(mockRepo, nil)

	tests := []struct {
		name       string
		dto        CreateUseCaseDTO
		mockDefine func()
		wantErr    bool
	}{
		{
			name: "success to create customer",
			dto: CreateUseCaseDTO{
				Name:       "test",
				Email:      "test@test.com",
				Prefecture: "tokyo",
				Address:    "shibuya",
				Password:   "password",
			},
			mockDefine: func() {
				mockRepo.On("Create", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "failed to create customer",
			dto: CreateUseCaseDTO{
				Name:       "test",
				Email:      "test@test.com",
				Prefecture: "tokyo",
				Address:    "shibuya",
				Password:   "password",
			},
			mockDefine: func() {
				mockRepo.On("Create", mock.Anything).Return(errors.New("failed to create customer"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.mockDefine()
			err := useCase.Execute(&gin.Context{}, tt.dto)
			if tt.wantErr && err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

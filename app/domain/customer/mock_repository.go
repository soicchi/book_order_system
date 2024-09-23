package customer

import (
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) Create(customer *Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

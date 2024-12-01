package orderdetail

// import (
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/mock"
// )

// type MockRepository struct {
// 	mock.Mock
// }

// func NewMockRepository() *MockRepository {
// 	return &MockRepository{}
// }

// func (m *MockRepository) BulkCreate(ctx echo.Context, orderDetails []*OrderDetail) error {
// 	args := m.Called(ctx, orderDetails)
// 	return args.Error(0)
// }

// func (m *MockRepository) FindByOrderID(ctx echo.Context, orderID string) (OrderDetails, error) {
// 	args := m.Called(ctx, orderID)
// 	if args.Get(0) == nil {
// 		return nil, args.Error(1)
// 	}

// 	return args.Get(0).(OrderDetails), args.Error(1)
// }

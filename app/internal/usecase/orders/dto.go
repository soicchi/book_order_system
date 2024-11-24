package orders

type CreateDetailInput struct {
	ProductID string
	Quantity  int
}

type CreateInput struct {
	CustomerID string
	OrderItems []*CreateDetailInput
}

func NewCreateInput(customerID string, orderItems []*CreateDetailInput) *CreateInput {
	return &CreateInput{
		CustomerID: customerID,
		OrderItems: orderItems,
	}
}

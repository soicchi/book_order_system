package orders

type CreateOrderRequest struct {
	CustomerID        string `param:"customer_id" validate:"required"`
	ShippingAddressID string `json:"shipping_address_id" validate:"required"`
}

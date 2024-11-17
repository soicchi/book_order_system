package shippingAddresses

type CreateShippingAddressRequest struct {
	Prefecture string `json:"prefecture" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	CustomerID string `param:"customer_id" validate:"required"`
}

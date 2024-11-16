package customers

import "github.com/soicchi/book_order_system/internal/usecase/dto"

type CustomerResponse struct {
	Customer *dto.CustomerOutput `json:"customer"`
	Message  string              `json:"message"`
}

func NewCustomerResponse(customer *dto.CustomerOutput, message string) *CustomerResponse {
	return &CustomerResponse{
		Customer: customer,
		Message:  message,
	}
}

package customers

type CreateCustomerRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type GetCustomerRequest struct {
	ID string `param:"id" validate:"required"`
}

package customers

type CreateCustomerRequest struct {
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Prefecture string `json:"prefecture" validate:"required"`
	Address    string `json:"address" validate:"required"`
	Password   string `json:"password" validate:"required"`
}

package customers

type CreateCustomerRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Prefecture string `json:"prefecture"`
	Address    string `json:"address"`
	Password   string `json:"password"`
}

package dto

type CreateCustomerInput struct {
	Name  string
	Email string
}

func NewCreateCustomerDTO(name, email string) *CreateCustomerInput {
	return &CreateCustomerInput{
		Name:  name,
		Email: email,
	}
}

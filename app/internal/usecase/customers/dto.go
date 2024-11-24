package customers

type CreateInput struct {
	Name  string
	Email string
}

func NewCreateInput(name, email string) *CreateInput {
	return &CreateInput{
		Name:  name,
		Email: email,
	}
}

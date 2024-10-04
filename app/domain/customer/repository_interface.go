package customer

type CustomerRepository interface {
	Create(customer *Customer) error
}

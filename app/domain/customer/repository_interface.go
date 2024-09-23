package customer

import ()

type CustomerRepository interface {
	Create(customer *Customer) error
}

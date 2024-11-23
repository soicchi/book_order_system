package payment

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/domain/values"

	"github.com/google/uuid"
)

type Payment struct {
	id        uuid.UUID
	amount    float64
	method    values.PaymentMethod
	status    values.PaymentStatus
	createdAt *time.Time
	updatedAt *time.Time
}

func New(amount float64, method, status string) (*Payment, error) {
	paymentUUID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("failed to generate uuid: %w", err)
	}

	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than 0: %f", amount)
	}

	paymentStatus, err := values.NewPaymentStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment status: %w", err)
	}

	paymentMethod, err := values.NewPaymentMethod(method)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}

	now := time.Now()

	return new(paymentUUID, amount, paymentMethod, paymentStatus, &now, &now), nil
}

func Reconstruct(
	id uuid.UUID,
	amount float64,
	method string,
	status string,
	createdAt *time.Time,
	updatedAt *time.Time,
) (*Payment, error) {
	paymentStatus, err := values.NewPaymentStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment status: %w", err)
	}

	paymentMethod, err := values.NewPaymentMethod(method)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment method: %w", err)
	}

	return new(id, amount, paymentMethod, paymentStatus, createdAt, updatedAt), nil
}

func new(
	id uuid.UUID,
	amount float64,
	method values.PaymentMethod,
	status values.PaymentStatus,
	createdAt *time.Time,
	updatedAt *time.Time,
) *Payment {
	return &Payment{
		id:        id,
		amount:    amount,
		method:    method,
		status:    status,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (p *Payment) ID() uuid.UUID {
	return p.id
}

func (p *Payment) Amount() float64 {
	return p.amount
}

func (p *Payment) Method() values.PaymentMethod {
	return p.method
}

func (p *Payment) Status() values.PaymentStatus {
	return p.status
}

func (p *Payment) CreatedAt() *time.Time {
	return p.createdAt
}

func (p *Payment) UpdatedAt() *time.Time {
	return p.updatedAt
}

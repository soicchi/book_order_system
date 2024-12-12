package fixtures

import (
	"time"

	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var TestTickets = map[string]models.Ticket{
	"ticket1": {
		ID:             uuid.New(),
		QRCode:         "ticket1_qr_code",
		Status:         "active",
		IssuedAt:       time.Now(),
		UsedAt:         time.Time{},
		RegistrationID: TestRegistrations["registration1"].ID,
	},
	"ticket2": {
		ID:             uuid.New(),
		QRCode:         "ticket2_qr_code",
		Status:         "active",
		IssuedAt:       time.Now(),
		UsedAt:         time.Time{},
		RegistrationID: TestRegistrations["registration2"].ID,
	},
}

func CreateTestTickets(db *gorm.DB) error {
	tickets := make([]models.Ticket, 0, len(TestTickets))
	for _, ticket := range TestTickets {
		tickets = append(tickets, ticket)
	}

	return db.Create(&tickets).Error
}

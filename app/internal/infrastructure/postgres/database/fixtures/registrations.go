package fixtures

import (
	"time"

	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var TestRegistrations = map[string]models.Registration{
	"registration1": {
		ID:           uuid.New(),
		RegisteredAt: time.Date(2024, time.December, 12, 10, 0, 0, 0, time.UTC),
		Status:       "registered",
		UserID:       TestUsers["attendee1"].ID,
		EventID:      TestEvents["event1"].ID,
	},
	"registration2": {
		ID:           uuid.New(),
		RegisteredAt: time.Date(2024, time.December, 12, 11, 0, 0, 0, time.UTC),
		Status:       "registered",
		UserID:       TestUsers["attendee1"].ID,
		EventID:      TestEvents["event2"].ID,
	},
}

func CreateTestRegistrations(db *gorm.DB) error {
	registrations := make([]models.Registration, 0, len(TestRegistrations))
	for _, registration := range TestRegistrations {
		registrations = append(registrations, registration)
	}

	return db.Create(&registrations).Error
}

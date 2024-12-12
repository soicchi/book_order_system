package fixtures

import (
	"time"

	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var TestEvents = map[string]models.Event{
	"event1": {
		ID:          uuid.New(),
		Title:       "event1_title",
		Description: "event1_description",
		StartDate:   time.Date(2024, time.December, 12, 10, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, time.December, 12, 10, 23, 59, 59, time.UTC),
		CreatedBy:   TestUsers["organizer1"].ID,
		VenueID:     TestVenues["venue1"].ID,
	},
	"event2": {
		ID:          uuid.New(),
		Title:       "event2_title",
		Description: "event2_description",
		StartDate:   time.Date(2024, time.December, 12, 11, 0, 0, 0, time.UTC),
		EndDate:     time.Date(2024, time.December, 12, 11, 23, 59, 59, time.UTC),
		CreatedBy:   TestUsers["organizer1"].ID,
		VenueID:     TestVenues["venue2"].ID,
	},
}

func CreateTestEvents(db *gorm.DB) error {
	events := make([]models.Event, 0, len(TestEvents))
	for _, event := range TestEvents {
		events = append(events, event)
	}

	return db.Create(&events).Error
}

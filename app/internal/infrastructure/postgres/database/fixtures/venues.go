package fixtures

import (
	"time"

	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var TestVenues = map[string]models.Venue{
	"venue1": {
		ID:        uuid.New(),
		Name:      "test_venue_1",
		Address:   "test_address_1",
		Capacity:  100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	"venue2": {
		ID:        uuid.New(),
		Name:      "test_venue_2",
		Address:   "test_address_2",
		Capacity:  200,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func CreateVenues(db *gorm.DB) error {
	venues := make([]models.Venue, 0, len(TestVenues))
	for _, venue := range TestVenues {
		venues = append(venues, venue)
	}

	return db.Create(&venues).Error
}

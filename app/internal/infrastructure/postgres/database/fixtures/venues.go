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
		Name:      "venue1_name",
		Address:   "venue1_address",
		Capacity:  100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	"venue2": {
		ID:        uuid.New(),
		Name:      "venue2_name",
		Address:   "venue2_address",
		Capacity:  200,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func CreateTestVenues(db *gorm.DB) error {
	venues := make([]models.Venue, 0, len(TestVenues))
	for _, venue := range TestVenues {
		venues = append(venues, venue)
	}

	return db.Create(&venues).Error
}

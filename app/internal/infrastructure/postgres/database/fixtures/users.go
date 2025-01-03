package fixtures

import (
	"time"

	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var TestUsers = map[string]models.User{
	"attendee1": {
		ID:        uuid.New(),
		Name:      "attendee1_user",
		Email:     "attendee1@test.com",
		Role:      "attendee",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	"organizer1": {
		ID:        uuid.New(),
		Name:      "organizer1_user",
		Email:     "organizer1@test.com",
		Role:      "organizer",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func CreateTestUsers(db *gorm.DB) error {
	users := make([]models.User, 0, len(TestUsers))
	for _, user := range TestUsers {
		users = append(users, user)
	}

	return db.Create(&users).Error
}

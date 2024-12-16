package repository

import (
	"testing"

	"event_system/internal/domain/role"
	"event_system/internal/domain/user"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestWithTransaction(t *testing.T) {
	tests := []struct {
		name        string
		user        *user.User
		anotherUser *user.User
		wantErr     bool
	}{
		{
			name: "Create user and venue successfully",
			user: user.New(
				"user1",
				"test@test.com",
				role.New(role.Attendee),
			),
			anotherUser: user.New(
				"user2",
				"anothertest@test.com",
				role.New(role.Attendee),
			),
			wantErr: false,
		},
		{
			name: "Subsequent process fails",
			user: user.New(
				"user1",
				"transaction_test@test.com",
				role.New(role.Attendee),
			),
			anotherUser: user.New(
				"user2",
				"transaction_test@test.com",
				role.New(role.Attendee),
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			db := database.GetDB(ctx)

			// delete created data in this test
			defer func() {
				users := []models.User{{ID: tt.user.ID()}, {ID: tt.anotherUser.ID()}}
				db.Delete(users)
			}()

			// Fetch users and venues before test
			var beforeUserModels []models.User
			if err := db.Find(&beforeUserModels).Error; err != nil {
				t.Fatalf("failed to fetch users: %v", err)
			}

			txManager := NewTransactionManager()
			userRepository := NewUserRepository()

			err := txManager.WithTransaction(ctx, func(ctx echo.Context) error {
				if err := userRepository.Create(ctx, tt.user); err != nil {
					return err
				}

				if err := userRepository.Create(ctx, tt.anotherUser); err != nil {
					return err
				}

				return nil
			})

			if tt.wantErr {
				assert.Error(t, err)

				// Check if the transaction is rolled back
				var afterUserModels []models.User
				if err = db.Find(&afterUserModels).Error; err != nil {
					t.Fatalf("failed to fetch users: %v", err)
				}

				var afterVenueModels []models.Venue
				if err = db.Find(&afterVenueModels).Error; err != nil {
					t.Fatalf("failed to fetch venues: %v", err)
				}

				assert.Equal(t, len(beforeUserModels), len(afterUserModels))
				return
			}

			assert.NoError(t, err)

			// Check if the transaction is committed
			var userModel models.User
			if err = db.Where("id = ?", tt.user.ID()).First(&userModel).Error; err != nil {
				t.Fatalf("failed to fetch user: %v", err)
			}

			assert.Equal(t, tt.user.ID(), userModel.ID)
			assert.Equal(t, tt.user.Name(), userModel.Name)
			assert.Equal(t, tt.user.Email(), userModel.Email)
			assert.Equal(t, tt.user.Role().Value().String(), userModel.Role)

			var anotherUserModel models.User
			if err = db.Where("id = ?", tt.anotherUser.ID()).First(&anotherUserModel).Error; err != nil {
				t.Fatalf("failed to fetch user: %v", err)
			}

			assert.Equal(t, tt.anotherUser.ID(), anotherUserModel.ID)
			assert.Equal(t, tt.anotherUser.Name(), anotherUserModel.Name)
			assert.Equal(t, tt.anotherUser.Email(), anotherUserModel.Email)
			assert.Equal(t, tt.anotherUser.Role().Value().String(), anotherUserModel.Role)
		})
	}
}

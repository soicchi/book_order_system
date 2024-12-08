package repository

import (
	"testing"
	"time"

	"event_system/internal/domain/role"
	"event_system/internal/domain/user"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/database/fixtures"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    *user.User
		wantErr bool
	}{
		{
			name: "Create user successfully",
			user: user.New(
				"user1",
				"user1@test.com",
				role.New(role.Attendee),
			),
			wantErr: false,
		},
		{
			name: "Create user with duplicated email",
			user: user.New(
				"user1",
				fixtures.TestUsers["attendee1"].Email, // Duplicated email
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

			// トランザクションを開始
			tx, err := database.BeginTx(ctx)
			if err != nil {
				t.Fatalf("Failed to begin transaction: %v", tx.Error)
			}

			// テスト終了時にロールバック
			defer func() {
				if err := tx.Rollback().Error; err != nil {
					t.Fatalf("Failed to rollback transaction: %v", err)
				}
			}()

			var beforeUserModels []models.User
			if err := db.Find(&beforeUserModels).Error; err != nil {
				t.Fatalf("failed to fetch users: %v", err)
			}

			r := NewUserRepository()

			repoErr := r.Create(ctx, tt.user)

			if tt.wantErr {
				assert.NotNil(t, repoErr)

				var afterUserModels []models.User

				if err = db.Find(&afterUserModels).Error; err != nil {
					t.Fatalf("failed to fetch user: %v", err)
				}
				assert.Equal(t, len(afterUserModels), len(beforeUserModels))
				return
			}

			assert.Nil(t, repoErr)

			var userModel models.User
			if err = tx.Where("id = ?", tt.user.ID()).First(&userModel).Error; err != nil {
				t.Fatalf("failed to fetch user: %v", err)
			}

			assert.Equal(t, tt.user.ID(), userModel.ID)
			assert.Equal(t, tt.user.Name(), userModel.Name)
			assert.Equal(t, tt.user.Email(), userModel.Email)
			assert.Equal(t, tt.user.Role().Value().String(), userModel.Role)
		})
	}
}

func TestFetchByID(t *testing.T) {
	tests := []struct {
		name    string
		id      uuid.UUID
		wantErr bool
	}{
		{
			name:    "Fetch user successfully",
			id:      fixtures.TestUsers["attendee1"].ID,
			wantErr: false,
		},
		{
			name:    "Fetch user not found",
			id:      uuid.New(),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewUserRepository()
			user, err := r.FetchByID(ctx, tt.id)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, user)
				return
			}

			// Does not exist
			if user == nil {
				assert.Nil(t, err)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, user)
		})
	}
}

func TestFetchAll(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Fetch all users successfully",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)

			r := NewUserRepository()
			users, err := r.FetchAll(ctx)

			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Nil(t, users)
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, users)
			assert.Equal(t, len(users), len(fixtures.TestUsers))
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name    string
		user    *user.User
		wantErr bool
	}{
		{
			name: "Update user successfully",
			user: user.Reconstruct(
				fixtures.TestUsers["attendee1"].ID,
				"updateduser1",
				"updated@test.com",
				role.New(role.Organizer),
				fixtures.TestUsers["attendee1"].CreatedAt,
				time.Now(),
			),
			wantErr: false,
		},
		{
			name: "Update user with duplicated email",
			user: user.Reconstruct(
				fixtures.TestUsers["attendee1"].ID,
				"updateduser1",
				fixtures.TestUsers["organizer1"].Email, // Duplicated email
				role.New(role.Organizer),
				fixtures.TestUsers["attendee1"].CreatedAt,
				time.Now(),
			),
			wantErr: true,
		},
		{
			name: "Update user not found",
			user: user.Reconstruct(
				uuid.New(),
				"updateduser1",
				"update@test.com",
				role.New(role.Organizer),
				fixtures.TestUsers["attendee1"].CreatedAt,
				time.Now(),
			),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			ctx := e.NewContext(nil, nil)
			db := database.GetDB(ctx)

			// トランザクションを開始
			tx, err := database.BeginTx(ctx)
			if err != nil {
				t.Fatalf("Failed to begin transaction: %v", tx.Error)
			}

			// テスト終了時にロールバック
			defer func() {
				if err := tx.Rollback().Error; err != nil {
					t.Fatalf("Failed to rollback transaction: %v", err)
				}
			}()

			r := NewUserRepository()

			repoErr := r.Update(ctx, tt.user)

			if tt.wantErr {
				assert.NotNil(t, repoErr)

				var userModel models.User
				if err = db.Where("id = ?", tt.user.ID()).First(&userModel).Error; err != nil {
					// Does not exist
					return
				}

				assert.NotEqual(t, tt.user.Name(), userModel.Name)
				assert.NotEqual(t, tt.user.Email(), userModel.Email)
				assert.NotEqual(t, tt.user.Role().Value().String(), userModel.Role)
				return
			}

			assert.Nil(t, repoErr)

			var userModel models.User
			if err = tx.Where("id = ?", tt.user.ID()).First(&userModel).Error; err != nil {
				t.Fatalf("failed to fetch user: %v", err)
			}

			assert.Equal(t, tt.user.ID(), userModel.ID)
			assert.Equal(t, tt.user.Name(), userModel.Name)
			assert.Equal(t, tt.user.Email(), userModel.Email)
			assert.Equal(t, tt.user.Role().Value().String(), userModel.Role)
		})
	}
}

package repository

import (
	"errors"
	"fmt"

	"event_system/internal/domain/role"
	"event_system/internal/domain/user"
	errs "event_system/internal/errors"
	"event_system/internal/infrastructure/postgres/database"
	"event_system/internal/infrastructure/postgres/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(ctx echo.Context, user *user.User) error {
	db := database.GetDB(ctx)

	userModel := models.User{
		ID:        user.ID(),
		Name:      user.Name(),
		Email:     user.Email(),
		Role:      user.Role().Value().String(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}

	err := db.Create(&userModel).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return errs.New(
			fmt.Errorf("user with email already exists: %w", err),
			errs.AlreadyExistError,
			errs.WithField("Email"),
		)
	}

	if err != nil {
		return errs.New(
			fmt.Errorf("failed to create user: %w", err),
			errs.UnexpectedError,
		)
	}

	return nil
}

func (r *UserRepository) FetchByID(ctx echo.Context, userID uuid.UUID) (*user.User, error) {
	db := database.GetDB(ctx)

	var userModel models.User
	err := db.Where("id = ?", userID).First(&userModel).Error
	// Return nil if user not found because of making use case more flexible
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch user: %w", err),
			errs.UnexpectedError,
		)
	}

	// Reconstruct user entity
	return user.Reconstruct(
		userModel.ID,
		userModel.Name,
		userModel.Email,
		role.Reconstruct(role.FromString(userModel.Role)),
		userModel.CreatedAt,
		userModel.UpdatedAt,
	), nil
}

func (r *UserRepository) FetchAll(ctx echo.Context) ([]*user.User, error) {
	db := database.GetDB(ctx)

	var userModels []models.User
	if err := db.Find(&userModels).Error; err != nil {
		return nil, errs.New(
			fmt.Errorf("failed to fetch users: %w", err),
			errs.UnexpectedError,
		)
	}

	// Reconstruct users entity
	users := make([]*user.User, 0, len(userModels))
	for _, userModel := range userModels {
		users = append(users, user.Reconstruct(
			userModel.ID,
			userModel.Name,
			userModel.Email,
			role.Reconstruct(role.FromString(userModel.Role)),
			userModel.CreatedAt,
			userModel.UpdatedAt,
		))
	}

	return users, nil
}

func (r *UserRepository) Update(ctx echo.Context, user *user.User) error {
	db := database.GetDB(ctx)

	userModel := models.User{
		ID:        user.ID(),
		Name:      user.Name(),
		Email:     user.Email(),
		Role:      user.Role().Value().String(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}

	err := db.Save(&userModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.New(
			fmt.Errorf("user not found: %w", err),
			errs.NotFoundError,
			errs.WithField("User"),
		)
	}

	if err != nil {
		return errs.New(
			fmt.Errorf("failed to update user: %w", err),
			errs.UnexpectedError,
		)
	}

	return nil
}

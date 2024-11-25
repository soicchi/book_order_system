package repository

import (
	"errors"
	"fmt"

	"github.com/soicchi/book_order_system/internal/domain/user"
	ers "github.com/soicchi/book_order_system/internal/errors"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/database"
	"github.com/soicchi/book_order_system/internal/infrastructure/postgres/models"

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

	err := db.Create(&models.User{
		ID:        user.ID(),
		Username:  user.Username(),
		Email:     user.Email(),
		Password:  user.Password().Value(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}).Error
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return ers.New(
			fmt.Errorf("user with email %s already exists", user.Email()),
			ers.AlreadyExist,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to create user: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}

func (r *UserRepository) FindByID(ctx echo.Context, id uuid.UUID) (*user.User, error) {
	db := database.GetDB(ctx)

	var u models.User
	err := db.Where("id = ?", id).First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find user: %w", err),
			ers.InternalServerError,
		)
	}

	return user.Reconstruct(
		u.ID,
		u.Username,
		u.Email,
		u.Password,
		u.CreatedAt,
		u.UpdatedAt,
	), nil
}

func (r *UserRepository) FindAll(ctx echo.Context) ([]*user.User, error) {
	db := database.GetDB(ctx)

	var users []*models.User
	err := db.Find(&users).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []*user.User{}, nil
	}

	if err != nil {
		return nil, ers.New(
			fmt.Errorf("failed to find users: %w", err),
			ers.InternalServerError,
		)
	}

	result := make([]*user.User, 0, len(users))
	for _, u := range users {
		result = append(result, user.Reconstruct(
			u.ID,
			u.Username,
			u.Email,
			u.Password,
			u.CreatedAt,
			u.UpdatedAt,
		))
	}

	return result, nil
}

func (r *UserRepository) Update(ctx echo.Context, user *user.User) error {
	db := database.GetDB(ctx)

	err := db.Model(&models.User{}).Where("id = ?", user.ID()).Updates(&models.User{
		Username: user.Username(),
		Email:    user.Email(),
		Password: user.Password().Value(),
	}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ers.New(
			fmt.Errorf("user with id %s not found", user.ID()),
			ers.NotFound,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to update user: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}

func (r *UserRepository) Delete(ctx echo.Context, id uuid.UUID) error {
	db := database.GetDB(ctx)

	err := db.Where("id = ?", id).Delete(&models.User{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ers.New(
			fmt.Errorf("user with id %s not found", id),
			ers.NotFound,
		)
	}

	if err != nil {
		return ers.New(
			fmt.Errorf("failed to delete user: %w", err),
			ers.InternalServerError,
		)
	}

	return nil
}

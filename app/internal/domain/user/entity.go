package user

import (
	"fmt"
	"time"

	"github.com/soicchi/book_order_system/internal/errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Userはユーザー情報を表すEntity
// 外部の層から直接変更されないようにするためにフィールドはprivateにする
// フィールドにアクセスするためのgetterメソッドを提供する
type User struct {
	id        uuid.UUID
	username  string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(username string, email string, password string) (*User, error) {
	// convert password to hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(
			fmt.Errorf("failed to generate hash from password: %w", err),
			errors.InternalServerError,
		)
	}

	return &User{
		id:        uuid.New(),
		username:  username,
		email:     email,
		password:  string(passwordHash),
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// DBからデータを取得した際にEntityに変換するための関数
func Reconstruct(id uuid.UUID, username, email, password string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		username:  username,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func new(id uuid.UUID, username, email, password string, createdAt, updatedAt time.Time) *User {
	return &User{
		id:        id,
		username:  username,
		email:     email,
		password:  password,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

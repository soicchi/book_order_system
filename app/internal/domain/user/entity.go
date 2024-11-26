package user

import (
	"time"

	"github.com/soicchi/book_order_system/internal/domain/values"

	"github.com/google/uuid"
)

// Userはユーザー情報を表すEntity
// 外部の層から直接変更されないようにするためにフィールドはprivateにする
// フィールドにアクセスするためのgetterメソッドを提供する
type User struct {
	id        uuid.UUID
	username  string
	email     string
	password  values.Password
	createdAt time.Time
	updatedAt time.Time
}

// Entityに関するビジネスルールに基づくバリデーションは初期化時のNew関数で行う
func New(username string, email string, password string) (*User, error) {
	passwordHash, err := values.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		id:        uuid.New(),
		username:  username,
		email:     email,
		password:  passwordHash,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// DBからデータを取得した際にEntityに変換するための関数
func Reconstruct(id uuid.UUID, username, email, password string, createdAt, updatedAt time.Time) *User {
	passwordHash := values.ReconstructPassword(password)

	return &User{
		id:        id,
		username:  username,
		email:     email,
		password:  passwordHash,
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

func (u *User) Password() values.Password {
	return u.password
}

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) Update(username, email, password string) error {
	passwordHash, err := values.NewPassword(password)
	if err != nil {
		return err
	}

	u.username = username
	u.email = email
	u.password = passwordHash
	u.updatedAt = time.Now()

	return nil
}

package domain

import (
	"context"
	"database/sql"
	"dione-backend/dto"
)

const (
	UserRoleAdmin string = "admin"
	UserRoleUser  string = "user"
)

type User struct {
	ID         string       `db:"id" json:"id"`
	Name       string       `db:"name" json:"name"`
	Email      string       `db:"email" json:"email"`
	Phone      string       `db:"phone" json:"phone"`
	Role       string       `db:"role" json:"role"`
	Bio        string       `db:"bio" json:"bio"`
	Password   string       `db:"password" json:"-"`
	AvatarPath string       `db:"avatar_path" json:"avatar_path"`
	IsActive   bool         `db:"is_active" json:"is_active"`
	CreatedAt  sql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at" json:"updated_at"`
}

type UserFilter struct {
	Limit     int64
	Cursor    string
	SortBy    string
	SortOrder string // ASC or DESC
}

type UserRepository interface {
	FindAll(ctx context.Context, filter UserFilter) ([]User, error)
	FindByID(ctx context.Context, id string) (User, error)
	FindByIDs(ctx context.Context, ids []string) (result []User, err error)
	FindByEmail(ctx context.Context, email string) (User, error)
	Save(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}

type UserService interface {
	Show(ctx context.Context, id string) (dto.UserData, error)
}

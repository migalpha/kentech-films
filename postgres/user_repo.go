package postgres

import (
	"fmt"
	"time"

	film "github.com/migalpha/kentech-films"

	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

func (repo UsersRepo) UserbyUsername(id film.Username) (film.User, error) {
	data := PostgresUser{
		Username: id.String(),
	}

	repo.DB.Where(&data).First(&data)
	user, err := repo.decodeUser(data)
	if err != nil {
		return film.User{}, fmt.Errorf("can't decode user: %w", err)
	}

	return user, nil
}

func (repo UsersRepo) UserByID(id film.UserID) (film.User, error) {
	data := PostgresUser{
		ID: id.Uint(),
	}

	repo.DB.Where(&data).First(&data)
	user, err := repo.decodeUser(data)
	if err != nil {
		return film.User{}, fmt.Errorf("can't decode user: %w", err)
	}

	return user, nil
}

func (repo UsersRepo) Save(user film.User) (film.UserID, error) {
	data := encodeUser(user)
	result := repo.DB.Create(&data)
	if result.Error != nil {
		return 0, fmt.Errorf("[UsersRepo:Save][err:%w]", result.Error)
	}
	return film.UserID(data.ID), nil
}

type PostgresUser struct {
	ID        uint      `json:"id,omitempty" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func encodeUser(user film.User) PostgresUser {
	return PostgresUser{
		Username:  user.Username.String(),
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		IsActive:  user.IsActive,
	}
}

func (repo UsersRepo) decodeUser(data PostgresUser) (film.User, error) {
	return film.User{
		ID:        film.UserID(data.ID),
		Username:  film.Username(data.Username),
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		IsActive:  data.IsActive,
	}, nil
}

// TableName overrides the table name used by PostgresUser to `users`
func (PostgresUser) TableName() string {
	return "users"
}

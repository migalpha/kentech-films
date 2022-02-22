package film

import (
	"context"
	"time"
)

type User struct {
	ID        UserID    `json:"id"`
	Username  Username  `json:"username"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type UserID uint
type Username string

func (username Username) String() string {
	return string(username)
}

func (id UserID) Uint() uint {
	return uint(id)
}

//go:generate mockery --name UserProvider
type UserProvider interface {
	UserbyUsername(context.Context, Username) (User, error)
	UserByID(context.Context, UserID) (User, error)
}

//go:generate mockery --name UserSaver
type UserSaver interface {
	Save(context.Context, User) (UserID, error)
}

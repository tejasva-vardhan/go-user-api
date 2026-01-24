package store

import (
	"context"

	"github.com/tejasva-vardhan/go-user-api/model"
)

// UserStore defines what operations our storage must support
// Handler ko farak nahi padega store in-memory hai ya DB based
type UserRepository interface {

	// CreateUser inserts a new user and returns created user with ID
	CreateUser(ctx context.Context, user model.User) (model.User, error)

	// GetAllUsers returns list of users
	GetAllUsers(ctx context.Context) ([]model.User, error)

	// GetUserByID returns single user by id
	GetUserByID(ctx context.Context, id int64) (model.User, error)

	// UpdateUser updates user by id
	UpdateUser(ctx context.Context, id int64, user model.User) (model.User, error)

	// DeleteUser deletes user by id
	DeleteUser(ctx context.Context, id int64) error
}
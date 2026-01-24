package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tejasva-vardhan/go-user-api/model"
)

// SQLiteUserStore implements UserRepository using SQLite DB
type SQLiteUserStore struct {
	DB *sql.DB
} // DB reference

// NewSQLiteUserStore creates new SQLite store
func NewSQLiteUserStore(db *sql.DB) *SQLiteUserStore {
	return &SQLiteUserStore{DB: db}
} // constructor

// CreateUser inserts a new user into DB
func (s *SQLiteUserStore) CreateUser(ctx context.Context, user model.User) (model.User, error) {

	// SQL query: insert name and email
	// id auto generate hoga
	query := `INSERT INTO users (name, email) VALUES (?, ?)`

	// ExecContext = query execute with context support
	result, err := s.DB.ExecContext(ctx, query, user.Name, user.Email)
	if err != nil {
		return model.User{}, fmt.Errorf("create user failed: %w", err)
	} // insert failed

	// LastInsertId = new user id
	id, err := result.LastInsertId()
	if err != nil {
		return model.User{}, fmt.Errorf("get last insert id failed: %w", err)
	} // id fetch failed

	user.ID = id // set generated id
	return user, nil
}

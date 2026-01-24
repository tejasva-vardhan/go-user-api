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
// GetAllUsers returns all users from DB
func (s *SQLiteUserStore) GetAllUsers(ctx context.Context) ([]model.User, error) {

	// Query: saare users fetch karo
	query := `SELECT id, name, email FROM users ORDER BY id DESC`

	// QueryContext = rows return karega
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("get all users failed: %w", err)
	} // query failed

	// rows close mandatory
	defer rows.Close() // cleanup

	users := make([]model.User, 0) // empty list

	// rows.Next() = each row iterate
	for rows.Next() {

		var u model.User // one user object

		// Scan = row values u me daalo
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, fmt.Errorf("scan user failed: %w", err)
		} // scan failed

		users = append(users, u) // add to list
	}

	// iteration ke baad check for errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	} // row loop error

	return users, nil
}

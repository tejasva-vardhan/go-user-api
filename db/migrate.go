package db

import (
	"database/sql"
	"fmt"
)

// Migrate creates required tables if they don't exist
func Migrate(database *sql.DB) error {

	// SQL query for users table
	// IF NOT EXISTS = already table hai to dobara create nahi karega
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	// Exec = query run karo (no rows return)
	_, err := database.Exec(query)
	if err != nil {
		return fmt.Errorf("users table migrate failed: %w", err)
	} // migration error

	return nil
}

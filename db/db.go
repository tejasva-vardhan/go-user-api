package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// OpenSQLite opens SQLite database connection
func OpenSQLite() (*sql.DB, error) {

	// DB_PATH env variable se path lene ki try
	// agar env empty ho to default path use hoga
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/app.db"
	} // default path set

	// sql.Open creates DB handle (connection nahi banata immediately)
	database, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite failed: %w", err)
	} // open failed

	// Ping ensures DB file + connection ok
	err = database.Ping()
	if err != nil {
		return nil, fmt.Errorf("sqlite ping failed: %w", err)
	} // ping failed

	return database, nil
}

package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	_ "github.com/tursodatabase/go-libsql"
	"github.com/tursodatabase/go-libsql/example/remote/envs"
)

func main() {
	setupSlog()
	if err := run(); err != nil {
		slog.Error("error running example", "error", err)
		os.Exit(1)
	}
}

func run() (err error) {
	err = envs.Load()
	if err != nil {
		return fmt.Errorf("error loading environment variables: %w", err)
	}
	// Get database URL and auth token from environment variables
	dbUrl := envs.TURSO_URL
	if dbUrl == "" {
		return fmt.Errorf("TURSO_URL environment variable not set")
	}

	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	if authToken != "" {
		dbUrl += "?authToken=" + authToken
	}

	// Open database connection
	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		return fmt.Errorf("error opening cloud db: %w", err)
	}
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			slog.Error("error closing libsql db", "error", closeErr)
			if err == nil {
				err = closeErr
			}
		} else {
			slog.Debug("closed libsql db")
		}
	}()
	// Configure connection pool: https://github.com/tursodatabase/go-libsql/issues/13
	db.SetConnMaxIdleTime(9 * time.Second)

	// Create test table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	// Check if test data already exists
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM test WHERE id = 1)").Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking existing data: %w", err)
	}

	// Insert test data only if it doesn't exist
	if !exists {
		_, err = db.Exec("INSERT INTO test (id, name) VALUES (?, ?)", 1, "remote test")
		if err != nil {
			return fmt.Errorf("error inserting data: %w", err)
		}
		slog.Debug("inserted test data")
	} else {
		slog.Debug("test data already exists, skipping insert")
	}

	// Query the data
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		return fmt.Errorf("error querying data: %w", err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			slog.Error("error closing rows", "error", closeErr)
			if err == nil {
				err = closeErr
			}
		}
	}()

	// Print results
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}
		fmt.Printf("Row: id=%d, name=%s\n", id, name)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("error iterating rows: %w", err)
	}

	slog.Debug("successfully connected and executed queries", "url", envs.TURSO_URL)
	return nil
}

func setupSlog() {
	// Configure slog to only show log level
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Remove time attribute
			if a.Key == "time" {
				return slog.Attr{}
			}
			return a
		},
	}))
	slog.SetDefault(logger)
}

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/tursodatabase/go-libsql"
	"os"
)

func run() (err error) {
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	db, err := sql.Open("libsql", "file:"+dir+"/test.db")
	if err != nil {
		return err
	}
	defer func() {
		if closeError := db.Close(); closeError != nil {
			fmt.Println("Error closing database", closeError)
			if err == nil {
				err = closeError
			}
		}
	}()

	_, err = db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name TEXT)")
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		_, err = db.Exec(fmt.Sprintf("INSERT INTO test (id, name) VALUES (%d, 'test-%d')", i, i))
		if err != nil {
			return err
		}
	}

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		return err
	}
	defer func() {
		if closeError := rows.Close(); closeError != nil {
			fmt.Println("Error closing rows", closeError)
			if err == nil {
				err = closeError
			}
		}
	}()
	i := 0
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return err
		}
		if id != i {
			return fmt.Errorf("expected id %d, got %d", i, id)
		}
		if name != fmt.Sprintf("test-%d", i) {
			return fmt.Errorf("expected name %s, got %s", fmt.Sprintf("test-%d", i), name)
		}
		i++
	}
	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

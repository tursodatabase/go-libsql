package main

import "database/sql"
import _ "github.com/libsql/go-libsql"

func main() {
	db, err := sql.Open("libsql", ":memory:")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
}

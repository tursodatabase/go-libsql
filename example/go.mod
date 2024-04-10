module github.com/tursodatabase/go-libsql/example

go 1.20

require github.com/tursodatabase/go-libsql v0.0.0-20230829151150-d30caadc3a7c

require (
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/libsql/sqlite-antlr4-parser v0.0.0-20240327125255-dbf53b6cbf06 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
)

replace github.com/tursodatabase/go-libsql => ../../go-libsql

module github.com/tursodatabase/go-libsql/example

go 1.20

require github.com/tursodatabase/go-libsql v0.0.0-20230829151150-d30caadc3a7c

require (
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230512164433-5d1fd1a340c9 // indirect
	github.com/libsql/sqlite-antlr4-parser v0.0.0-20230802215326-5cb5bb604475 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
)

replace github.com/tursodatabase/go-libsql => ../../go-libsql

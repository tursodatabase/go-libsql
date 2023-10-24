# LibSQL package for Go

[libSQL](https://github.com/libsql/libsql) is an open source, open contribution fork of SQLite.
This source repository contains libSQL API bindings for Go.

## Notice
This package comes with a precompiled native libraries.
Currently only `darwin amd64` and `darwin arm64` are supported.
We're working on adding support for more platforms.

## Features

* In-memory databases and local database files, like SQLite
* Remote database access to libSQL server
* In-app replica that syncs with a libSQL server

## Installing

```
go get github.com/libsql/go-libsql
```

`go-libsql` uses `CGO` to make calls to LibSQL. You must build your binaries with `CGO_ENABLED=1`.

## Getting Started

### Connecting to the database

To connect to the database one needs to create a `libsql.Connector` using one of the factory functions: `libsql.NewEmbeddedReplicaConnector` or `libsql.NewEmbeddedReplicaConnectorWithAutoSync`.

Here's an example of obtaining a `sql.DB` object from `database/sql` package:

```
dbPath := // Path do db file on local disk
primaryUrl := // URL to primary database instance
connector := NewEmbeddedReplicaConnector(dbPath, primaryUrl, authToken)
db := sql.OpenDB(connector)
defer db.Close()
```

Once `sql.DB` object is created one can use it as any other database that supports `database/sql` package.

### Fetching updates from primary database instance

If the connector is created with `libsql.NewEmbeddedReplicaConnectorWithAutoSync` then it will automatically fetch updates from a primary periodically.

For connectors created with `libsql.NewEmbeddedReplicaConnector` we need to fetch updates manually by calling `connector.Sync`

## Examples

Module with usage examples can be found in [example directory].

## License

This project is licensed under the [MIT license].

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in libSQL by you, shall be licensed as MIT, without any additional
terms or conditions.

[MIT license]: https://github.com/libsql/go-libsql/blob/main/LICENSE
[example directory]: https://github.com/libsql/go-libsql/tree/main/example

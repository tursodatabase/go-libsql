<p align="center">
  <a href="https://tur.so/turso-go">
    <picture>
      <img src="/.github/cover.png" alt="libSQL Go" />
    </picture>
  </a>
  <h1 align="center">libSQL Go</h1>
</p>

<p align="center">
  Databases for Go multi-tenant AI Apps.
</p>

<p align="center">
  <a href="https://tur.so/turso-go"><strong>Turso</strong></a> ·
  <a href="https://docs.turso.tech"><strong>Docs</strong></a> ·
  <a href="https://docs.turso.tech/sdk/go/quickstart"><strong>Quickstart</strong></a> ·
  <a href="https://docs.turso.tech/sdk/go/reference"><strong>SDK Reference</strong></a> ·
  <a href="https://turso.tech/blog"><strong>Blog &amp; Tutorials</strong></a>
</p>

<p align="center">
  <a href="LICENSE">
    <picture>
      <img src="https://img.shields.io/github/license/tursodatabase/go-libsql?color=0F624B" alt="MIT License" />
    </picture>
  </a>
  <a href="https://tur.so/discord-go">
    <picture>
      <img src="https://img.shields.io/discord/933071162680958986?color=0F624B" alt="Discord" />
    </picture>
  </a>
  <a href="#contributors">
    <picture>
      <img src="https://img.shields.io/github/contributors/tursodatabase/go-libsql?color=0F624B" alt="Contributors" />
    </picture>
  </a>
  <a href="/examples">
    <picture>
      <img src="https://img.shields.io/badge/browse-examples-0F624B" alt="Examples" />
    </picture>
  </a>
</p>

## Features

- 🔌 Works offline with [Embedded Replicas](https://docs.turso.tech/features/embedded-replicas/introduction)
- 🌎 Works with remote Turso databases
- ✨ Works with Turso [AI & Vector Search](https://docs.turso.tech/features/ai-and-embeddings)
- 🐘 Works Go PDO

> [!WARNING]
> This SDK is currently in technical preview. <a href="https://tur.so/discord-go">Join us in Discord</a> to report any issues.

## Install

```bash
go get github.com/tursodatabase/go-libsql
```

> [!NOTE]
> If you require a remote only, no-CGO Go libSQL driver, see [`libsql-client-go`](https://github.com/tursodatabase/libsql-client-go).

> [!NOTE]
> Currently only `linux amd64`, `linux arm64`, `darwin amd64` and `darwin arm64` are supported.

## Quickstart

The example below uses Embedded Replicas and syncs data every 1000ms from Turso.

```go
package main

import (
  "database/sql"
  "fmt"
  "os"
  "path/filepath"

  "github.com/tursodatabase/go-libsql"
)

func main() {
    dbName := "local.db"
    primaryUrl := "TURSO_DATABASE_URL"
    authToken := "TURSO_AUTH_TOKEN"
    syncInterval := time.Minute

    dir, err := os.MkdirTemp("", "libsql-*")
    if err != nil {
        fmt.Println("Error creating temporary directory:", err)
        os.Exit(1)
    }
    defer os.RemoveAll(dir)

    dbPath := filepath.Join(dir, dbName)

    connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl,
        libsql.WithAuthToken(authToken),
        libsql.WithSyncInterval(syncInterval),
    )
    if err != nil {
        fmt.Println("Error creating connector:", err)
        os.Exit(1)
    }
    defer connector.Close()

    db := sql.OpenDB(connector)
    defer db.Close()
}
```

## Documentation

Visit our [official documentation](https://docs.turso.tech/sdk/go).

## Support

Join us [on Discord](https://tur.so/discord-go) to get help using this SDK. Report security issues [via email](mailto:security@turso.tech).

## Contributors

See the [contributing guide](CONTRIBUTING.md) to learn how to get involved.

![Contributors](https://contrib.nn.ci/api?repo=tursodatabase/go-libsql)

<a href="https://github.com/tursodatabase/go-libsql/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22">
  <picture>
    <img src="https://img.shields.io/github/issues-search/tursodatabase/go-libsql?label=good%20first%20issue&query=label%3A%22good%20first%20issue%22%20&color=0F624B" alt="good first issue" />
  </picture>
</a>

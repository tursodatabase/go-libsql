# Run

```
go run main.go
```

- Environment variables are loaded from `envs/.env.local` by default.
- An example `.env.local` will be created if it doesn't exist.
- `local` mode will save a local.db file in this directory.
- `staging` and `prod` can be setup by creating `envs/.env.staging` and `envs/.env.prod` and setting the values from your Turso instance.
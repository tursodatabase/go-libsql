package envs

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

//go:embed .env.*
var fs embed.FS

// Environment Variables
var (
	TURSO_URL        string
	TURSO_AUTH_TOKEN string
	REMOTE_ENV       Env
)

type Env string

const (
	EnvLocal   Env = "local"
	EnvStaging Env = "staging"
	EnvProd    Env = "prod"
)

func (e Env) String() string {
	return string(e)
}

type EnvFile string

func (e Env) EnvFile() EnvFile {
	return EnvFile(".env." + e.String())
}

func (e EnvFile) Load() error {
	file, err := fs.Open(string(e))
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}
	return scanner.Err()
}

var didLoad = false

func Load() error {
	if didLoad {
		return nil
	}

	env, ok := os.LookupEnv("REMOTE_ENV")
	if !ok {
		slog.Info("REMOTE_ENV not set, using local")
		env = "local"
	}
	REMOTE_ENV = Env(env)

	eFile := REMOTE_ENV.EnvFile()

	// Check if .env.local exists, if not create it with default values
	if REMOTE_ENV == EnvLocal {
		if _, err := fs.Open(".env.local"); errors.Is(err, os.ErrNotExist) {
			slog.Info("Creating default .env.local file")
			defaultEnv := []byte("TURSO_URL=file:./local.db\nTURSO_AUTH_TOKEN=your-auth-token")
			if err := os.WriteFile("envs/.env.local", defaultEnv, 0644); err != nil {
				return fmt.Errorf("failed to create default .env.local: %w", err)
			}
			slog.Info("Created default .env.local file. The defaults will work for local testing, but you'll need to update the variables to connect to your Turso instance")
			slog.Info("Run `go run .` to start the application")
			os.Exit(1)
		}
	}

	slog.Info("Loading environment variables", "file", eFile)
	if err := eFile.Load(); err != nil {
		return err
	}

	envs := []envLookup{
		{&TURSO_URL, "TURSO_URL"},
		{&TURSO_AUTH_TOKEN, "TURSO_AUTH_TOKEN"},
	}
	if err := lookupEnvs(envs); err != nil {
		return err
	}

	didLoad = true
	slog.Debug("Loaded environment variables",
		"REMOTE_ENV", REMOTE_ENV,
		"DB_URL", TURSO_URL,
	)
	return nil
}

type envLookup struct {
	ptr *string
	key string
}

func lookupEnvs(envs []envLookup) error {
	var errorSlice []error
	for _, env := range envs {
		val, err := lookupEnv(env.key)
		if err != nil {
			errorSlice = append(errorSlice, err)
		}
		*env.ptr = val
	}
	if len(errorSlice) > 0 {
		return fmt.Errorf("failed to lookup envs: %w", errors.Join(errorSlice...))
	}
	return nil
}

func lookupEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("env %s not set", key)
	}
	return val, nil
}

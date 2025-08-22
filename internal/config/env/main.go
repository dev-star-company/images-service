package env

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

var (
	CLOUDFLARE_ACCOUNT_ID   string
	CLOUDFLARE_API_TOKEN    string
	CLOUDFLARE_DELIVERY_URL string
)

var envFile map[string]string

func init() {
	var err error

	if strings.ToLower(os.Getenv("LAUNCH_MODE")) == "debug" {
		absPath, pathErr := filepath.Abs(os.Args[0] + "/../.env")
		if pathErr != nil {
			panic("error resolving absolute path: " + pathErr.Error())
		}

		envFile, err = godotenv.Read(absPath)
	} else {
		envFile, err = godotenv.Read(".env")
	}
	if err != nil {
		panic("error reading .env file: " + err.Error())
	}

	CLOUDFLARE_ACCOUNT_ID = EnvToString("CLOUDFLARE_ACCOUNT_ID")
	CLOUDFLARE_API_TOKEN = EnvToString("CLOUDFLARE_API_TOKEN")
	CLOUDFLARE_DELIVERY_URL = EnvToString("CLOUDFLARE_DELIVERY_URL")
}

func ValidateEnv() error {
	requiredVars := []string{
		"CLOUDFLARE_ACCOUNT_ID", "CLOUDFLARE_API_TOKEN", "CLOUDFLARE_DELIVERY_URL",
	}
	for _, key := range requiredVars {
		if envFile[key] == "" {
			return errors.New(key + " is not set in .env file")
		}
	}
	return nil
}

func EnvToString(v string) string {
	return envFile[v]
}

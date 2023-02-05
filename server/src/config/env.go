package config

import (
	"fmt"
	"os"
)

var (
	PORT string

	ENVIRONMENT string

	POSTGRES_URL string

	Google_ID           string
	Google_SECRET       string
	GOOGLE_API_CLIENT   string
	GOOGLE_CALLBACK_API string

	CLIENT_DOMAIN string
)

func GetEnvAll() error {
	var err error
	PORT = GetEnvOrDefault("PORT", "8080")
	ENVIRONMENT = GetEnvOrDefault("ENVIRONMENT", "DEV")
	if IsPrd() {
		POSTGRES_URL, err = GetEssentialEnv("POSTGRES_URL")
	} else if IsDev() {
		POSTGRES_URL, err = DSN()
	} else {
		return fmt.Errorf("POSTGRES_URL is empty")
	}
	if err != nil {
		return err
	}
	Google_ID, err = GetEssentialEnv("Google_ID")
	if err != nil {
		return err
	}
	Google_SECRET, err = GetEssentialEnv("Google_SECRET")
	if err != nil {
		return err
	}
	GOOGLE_API_CLIENT, err = GetEssentialEnv("GOOGLE_API_CLIENT")
	if err != nil {
		return err
	}
	GOOGLE_CALLBACK_API, err = GetEssentialEnv("GOOGLE_CALLBACK_API")
	if err != nil {
		return err
	}
	CLIENT_DOMAIN, err = GetEssentialEnv("CLIENT_DOMAIN")
	if err != nil {
		return err
	}
	return nil
}

// GetEnvOrDefaultはenvPathに指定された環境変数を取得する
// 取得できなかった場合はdefaultEnvを利用する
func GetEnvOrDefault(envPath string, defaultEnv string) string {
	env := os.Getenv(envPath)
	if env == "" {
		return defaultEnv
	}
	return env
}

func GetEssentialEnv(envPath string) (string, error) {
	env := os.Getenv(envPath)
	if env == "" {
		return "", fmt.Errorf("essential env not found")
	}
	return env, nil
}

func IsDev() bool {
	return os.Getenv("ENV") == "DEV"
}

func IsPrd() bool {
	return os.Getenv("PRD") == "PRD"
}

package env

import (
	"fmt"
	"os"
)

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

func IsLocal() bool {
	return os.Getenv("ENV") == "DEV"
}

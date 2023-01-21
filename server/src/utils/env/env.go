package env

import (
	"log"
	"os"
)

// GetEnvOrDefaultはenvPathに指定された環境変数を取得する
// 取得できなかった場合はdefaultEnvを利用する
func GetEnvOrDefault(envPath string, defaultEnv string) string {
	env := os.Getenv(envPath)
	if env == "" {
		if defaultEnv == "" {
			log.Println("error: env not found and default env is empty too")
		}
		return defaultEnv
	}
	return env
}

func IsLocal() bool {
	return os.Getenv("ENV") == "DEV"
}

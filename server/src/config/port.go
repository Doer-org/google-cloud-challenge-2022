package config

import (
	"github.com/Doer-org/google-cloud-challenge-2022/utils/env"
)

// Portはサーバーのport番号を返します
func Port() string {
	return env.GetEnvOrDefault("PORT", "8080")
}

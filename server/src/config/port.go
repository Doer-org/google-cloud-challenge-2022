package config

import "github.com/Doer-org/google-cloud-challenge-2022/utils/helper"

// Portはサーバーのport番号を返します
func Port() string {
	return helper.GetEnvOrDefault("PORT", "8080")
}

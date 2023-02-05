package config

import (
	"fmt"
)

var (
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
	PGPASSWORD        string
	TZ                string
)

// DSNはdataSourceNameを返します、もし必須の環境変数が設定されてなかった場合はerrorを返します
func DSN() (string, error) {
	var err error
	POSTGRES_HOST, err = GetEssentialEnv("POSTGRES_HOST")
	if err != nil {
		return "", err
	}
	POSTGRES_PORT = GetEnvOrDefault("POSTGRES_PORT", "5432")
	POSTGRES_USER, err = GetEssentialEnv("POSTGRES_USER")
	if err != nil {
		return "", err
	}
	POSTGRES_PASSWORD, err = GetEssentialEnv("POSTGRES_PASSWORD")
	if err != nil {
		return "", err
	}
	POSTGRES_DB, err = GetEssentialEnv("POSTGRES_DB")
	if err != nil {
		return "", err
	}
	PGPASSWORD, err = GetEssentialEnv("PGPASSWORD")
	if err != nil {
		return "", err
	}
	TZ = GetEnvOrDefault("TZ", "Asia/Tokyo")

	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		POSTGRES_USER,
		POSTGRES_PASSWORD,
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_DB,
	), nil
}

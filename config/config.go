package config

import (
	"os"
	"strconv"
)

type VkConfig struct {
	CommunityAccessToken string
	CommunityID          string
	APIMethodPath        string
	APIVersion           string
}

func NewVkConfigFromEnv() *VkConfig {
	return &VkConfig{
		CommunityAccessToken: getEnv("COMMUNITY_ACCESS_TOKEN", ""),
		CommunityID:          getEnv("COMMUNITY_ID", ""),
		APIMethodPath:        getEnv("API_METHOD_PATH", "https://api.vk.com/method/"),
		APIVersion:           getEnv("API_VERSION", "5.122"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int32) int32 {
	vStr := getEnv(key, "")
	if v, err := strconv.Atoi(vStr); err == nil {
		return int32(v)
	}
	return defaultValue
}

func getEnvFloat(key string, defaultValue float32) float32 {
	vStr := getEnv(key, "")
	if v, err := strconv.ParseFloat(vStr, 32); err == nil {
		return float32(v)
	}
	return defaultValue
}

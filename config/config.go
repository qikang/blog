package config

import (
	"os"
)

type Config struct {
	Port         string
	PostsDir     string
	StaticDir    string
	TemplateDir  string
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", ":8083"),
		PostsDir:     getEnv("POSTS_DIR", "./posts"),
		StaticDir:    getEnv("STATIC_DIR", "./static"),
		TemplateDir:  getEnv("TEMPLATE_DIR", "./templates"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

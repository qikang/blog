package config

import (
	"os"
)

// Config 运行时配置。
// 注意: templates/static 已通过 go:embed 编译进二进制,
//这里不再持有 TemplateDir/StaticDir,它们随二进制分发。
type Config struct {
	Port     string
	PostsDir string
}

func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", ":8083"),
		PostsDir: getEnv("POSTS_DIR", "./posts"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

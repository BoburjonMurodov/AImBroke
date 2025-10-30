package config

import "os"

type Config struct {
	BotToken    string
	GeminiToken string
}

func Load() *Config {
	return &Config{
		BotToken:    getEnv("BOT_TOKEN", "YOUR_BOT_TOKEN"),
		GeminiToken: getEnv("GEMINI_TOKEN", "YOUR_GEMINI_TOKEN"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

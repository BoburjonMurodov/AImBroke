package config

import "os"

type Config struct {
	BotToken    string
	GeminiToken string
}

func Load() *Config {
	return &Config{
		BotToken:    os.Getenv("BOT_TOKEN"),
		GeminiToken: os.Getenv("GEMINI_TOKEN"),
	}
}

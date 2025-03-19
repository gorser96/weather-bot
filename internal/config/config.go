package config

import "os"

type Config struct {
	TelegramBotToken  string
	OpenWeatherAPIKey string
	WeatherapiAPIKey  string
}

func Load() *Config {
	return &Config{
		TelegramBotToken:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		OpenWeatherAPIKey: os.Getenv("OPENWEATHER_API_KEY"),
		WeatherapiAPIKey:  os.Getenv("WEATHERAPI_API_KEY"),
	}
}

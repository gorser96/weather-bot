package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"weather-bot/internal/config"
	"weather-bot/internal/handlers"
	"weather-bot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := config.Load()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	// Контекст для управления жизненным циклом программы
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ows := services.OpenWeatherService{APIKey: cfg.OpenWeatherAPIKey}
	ws := services.WeatherapiService{APIKey: cfg.WeatherapiAPIKey}
	oms := services.OpenmeteoService{OpenWeatherApiKey: cfg.OpenWeatherAPIKey}

	handlers.HandleTelegramCommands(bot, ctx, ows, ws, oms)
	handlers.StartScheduleHandle(bot, ctx, ows, ws, oms)

	// Обработка сигналов завершения
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнал завершения
	<-sigCh
	log.Println("Получен сигнал завершения...")

	// Отменяем контекст, чтобы завершить горутины
	cancel()

	// Ждем завершения всех горутин
	time.Sleep(1 * time.Second) // Даем время для завершения горутин
	log.Println("Бот завершил работу.")
}

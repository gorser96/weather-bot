package handlers

import (
	"context"
	"log"
	"time"
	"weather-bot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartScheduleHandle(bot *tgbotapi.BotAPI, ctx context.Context, weatherServices ...services.WeatherService) {
	go func() {
		ticker := time.NewTicker(10 * time.Second) // Отправляем каждые 10 секунд
		defer ticker.Stop()
		var lastSendedDay int = 0
		for {
			select {
			case <-ctx.Done():
				log.Println("Завершение отправки запланированных сообщений...")
				return
			case <-ticker.C:
				now := time.Now()
				nextSendTime := time.Date(
					now.Year(), now.Month(), now.Day(),
					7, 30, 0, 0, time.Local, // Запланированное время (7:30)
				)

				// Если время уже прошло сегодня, планируем на завтра
				if now.After(nextSendTime) && lastSendedDay != now.Day() {
					lastSendedDay = now.Day()

					chatList, err := services.ListChatTopics()
					if err != nil {
						log.Printf("%v", err)
						continue
					}

					var scheduledMsg = ""
					for _, ws := range weatherServices {
						weather, err := ws.GetDailyWeather("Stavropol")
						if err != nil {
							weather = err.Error()
						}
						scheduledMsg += weather + "\n"
					}

					for _, chatItem := range chatList.Chats {
						// Создание конфигурации сообщения
						msg := tgbotapi.NewMessage(chatItem.ChatID, scheduledMsg)
						msg.ReplyToMessageID = chatItem.TopicID
						_, err := bot.Send(msg)
						if err != nil {
							log.Printf("Ошибка при отправке сообщения: %v", err)
						} else {
							log.Println("Запланированное сообщение успешно отправлено!")
						}
					}
				}
			}
		}
	}()
}

package handlers

import (
	"context"
	"log"
	"strings"
	"weather-bot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const helpMsg = `Доступные команды:
/help, /? - справка
/start - отправлять ежедневно погоду в 7:30 в чат, откуда пришло сообщение
/stop - перестать отправлять ежедневно погоду в чат, откуда пришло сообшение
/weather - получить погоду на день в Ставрополе
/current - получить текущую погоду в Ставрополе`

func HandleTelegramCommands(bot *tgbotapi.BotAPI, ctx context.Context, weatherServices ...services.WeatherService) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Завершение обработки команд...")
				return
			case update := <-updates:
				if update.Message != nil {
					log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

					userInput := strings.Split(update.Message.Text, "@")[0]
					switch userInput {
					case "/?", "/help":
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMsg)
						msg.ReplyToMessageID = update.Message.ReplyToMessage.MessageID
						bot.Send(msg)

					case "/start":
						services.AddChatTopic(update.Message.Chat.ID, update.Message.ReplyToMessage.MessageID)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Теперь в этот чат будет отправляться погода")
						msg.ReplyToMessageID = update.Message.ReplyToMessage.MessageID
						bot.Send(msg)

					case "/stop":
						services.RemoveChatTopic(update.Message.Chat.ID, update.Message.ReplyToMessage.MessageID)
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Погода в этот чат больше не будет отправляться")
						msg.ReplyToMessageID = update.Message.ReplyToMessage.MessageID
						bot.Send(msg)

					case "/weather":
						var responseMsg = ""
						for _, ws := range weatherServices {
							weather, err := ws.GetDailyWeather("Stavropol")
							if err != nil {
								weather = err.Error()
							}
							responseMsg += weather + "\n"
						}

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseMsg)
						msg.ReplyToMessageID = update.Message.ReplyToMessage.MessageID
						bot.Send(msg)

					case "/current":
						var responseMsg = ""
						for _, ws := range weatherServices {
							weather, err := ws.GetCurrentWeather("Stavropol")
							if err != nil {
								weather = err.Error()
							}
							responseMsg += weather + "\n"
						}

						msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseMsg)
						msg.ReplyToMessageID = update.Message.ReplyToMessage.MessageID
						bot.Send(msg)

					default:
						break
					}
				}
			}
		}
	}()
}

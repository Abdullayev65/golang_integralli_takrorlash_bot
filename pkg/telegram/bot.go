package telegram

import (
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	service := service.NewService(bot)

	for update := range updates {
		if update.Message != nil {
			service.GotReq(update)
		}
	}
}

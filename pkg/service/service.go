package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Service struct {
	Bot *tgbotapi.BotAPI
}

func (s *Service) GotReq(update tgbotapi.Update) {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "ishlavotti : "+update.Message.Text)
	s.Bot.Send(message)
}

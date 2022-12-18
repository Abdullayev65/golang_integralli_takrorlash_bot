package service

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"time"
)

type Service struct {
	Bot *tgbotapi.BotAPI
}

func NewService(bot *tgbotapi.BotAPI) *Service {
	return &Service{Bot: bot}
}

func (s *Service) GotReq(update tgbotapi.Update) {
	text := update.Message.Text
	if len(text) > 0 && text[0] == '/' {
		s.commands(&update)
	} else {
		s.saveToSchedule(&update)
	}
}

func (s *Service) commands(update *tgbotapi.Update) {
	text := update.Message.Text
	var aswer string
	switch text {
	case "/help":
		aswer = "coming up..."
	case "/list":
		aswer = "coming up..."
	default:
		if strings.HasPrefix(text, "/delete") {
			aswer = "coming up..."
		} else {
			aswer = "no such command"
		}
	}
	s.send(update.Message.Chat.ID, aswer)
}

func (s *Service) saveToSchedule(update *tgbotapi.Update) {
	chatId := update.Message.Chat.ID
	message := tgbotapi.NewMessage(chatId, "ishlavotti : "+update.Message.Text)

	forward := tgbotapi.NewForward(chatId, chatId, update.Message.MessageID)
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println(s.Bot.Send(forward))
	}()
	s.Bot.Send(message)

}

func (s *Service) send(chatId int64, text string) {
	message := tgbotapi.NewMessage(chatId, text)
	s.Bot.Send(message)
}

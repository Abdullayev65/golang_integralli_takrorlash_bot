package service

import (
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/entity"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
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
	chatID, messageID := update.Message.Chat.ID, update.Message.MessageID
	data := entity.NewData(chatID, messageID)
	repository.SaveData(data)
	messageToSend := entity.NewMessageToSend(data)
	repository.SaveMessageToSend(messageToSend)
	s.sendReply(chatID, "saved \nID["+strconv.Itoa(data.Id)+"]", messageID)
}

func (s *Service) send(chatID int64, text string) {
	message := tgbotapi.NewMessage(chatID, text)
	s.Bot.Send(message)
}
func (s *Service) sendReply(chatId int64, text string, replyToMessageID int) {
	message := tgbotapi.NewMessage(chatId, text)
	message.ReplyToMessageID = replyToMessageID
	s.Bot.Send(message)
}

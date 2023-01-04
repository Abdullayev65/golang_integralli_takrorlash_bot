package service

import (
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/entity"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
	"strings"
)

var (
	exampleOfUpdateK   = "/k newK in id \n/k 1.6 in 37 - 1.6 koifsentni 37 id dagi xabarga joylash \n0.3 > k > 10; id xabar saqlanganda berilgan"
	comHelp            = readFile("textes/command-help")
	comStart           = readFile("textes/command-start")
	comAdvises         = readFile("textes/command-advises")
	advisesPhotoFIleID = tgbotapi.FileID("AgACAgIAAxkBAAOrY7PMeJetll91zRC9obfEnbZXLmAAAm_HMRuNiaBJppgJil1ij18BAAMCAANzAAMtBA")
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
	chatID := update.Message.Chat.ID
	switch text {
	case "/help":
		s.send(chatID, comHelp)
	case "/start":
		s.send(chatID, comStart)
	case "/advises":
		{
			photoCon := tgbotapi.NewPhoto(chatID, advisesPhotoFIleID)
			photoCon.Caption = comAdvises
			s.Bot.Send(photoCon)
		}
	case "/list":
		s.send(chatID, "coming up...")
	case "/integralli_takrorlash_haqida":
		{
			fileID := tgbotapi.FileID("BAACAgIAAxkBAAMHY7Kxl77r10_CGbidhwrp3NepVWMAAvsFAAIdVTBK0rqANq9sJ14tBA")
			videoConfig := tgbotapi.NewVideo(chatID, fileID)
			videoConfig.Caption = "Maskur video intervalli takrorlash haqida va YouTobe ning @Emuallim kanalidan olingan"
			s.Bot.Send(videoConfig)
		}
	default:
		switch {
		case strings.HasPrefix(text, "/k"):
			s.updateK(text, chatID, update.Message.MessageID)
		default:
			s.send(chatID, "no such command")
		}
	}
}

func (s *Service) saveToSchedule(update *tgbotapi.Update) {
	chatID, messageID := update.Message.Chat.ID, update.Message.MessageID
	data := entity.NewData(chatID, messageID)
	repository.SaveData(data)
	messageToSend := entity.NewMessageToSend(data)
	repository.SaveMessageToSend(messageToSend)
	s.sendReply(chatID, "saved["+strconv.Itoa(data.MessageID)+"]", messageID)
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

func (s *Service) updateK(text string, chatID int64, messageID int) {
	fields := strings.Fields(text)
	if len(fields) != 4 || fields[0] != "/k" || fields[2] != "in" {
		s.sendReply(chatID, "failed\n"+exampleOfUpdateK, messageID)
		return
	}
	newK, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		s.sendReply(chatID, fields[1]+" - k bo'lolmaydi\n1.7 - k bo'laoladi", messageID)
		return
	}
	if len(fields[1]) > 5 {
		s.sendReply(chatID, fields[1]+" - juda uzun son", messageID)
		return
	}
	if newK >= 10 || newK <= 0.3 {
		s.sendReply(chatID, fields[1]+" - 10 dan katta yoki 0.3 dan kichkina\n 0.3 > k > 10", messageID)
		return
	}
	messageId, err := strconv.Atoi(fields[3])
	if err != nil {
		s.sendReply(chatID, fields[3]+" - id emas", messageID)
		return
	}
	dataID, err := repository.GetIdOfData(messageId, chatID)
	if err != nil {
		s.sendReply(chatID, "id is gone wrong", messageID)
		return
	}
	err = repository.UpdateK(dataID, newK)
	if err != nil {
		s.sendReply(chatID, "sorry :( \nPlease try it later", messageID)
		return
	}
	s.sendReply(chatID, "successfully updated", messageID)
}
func readFile(name string) string {
	file, err := os.ReadFile(name)
	if err != nil {
		panic(name + " - file not founded")
	}
	return string(file)
}

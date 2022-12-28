package schedule

import (
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/entity"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

type Schedule struct {
	Bot *tgbotapi.BotAPI
}

const (
	halfHour    = int64(30 * time.Minute)
	tenMinutesD = 10 * time.Minute
	tenMinutes  = int64(tenMinutesD)
)

func NewSchedule(bot *tgbotapi.BotAPI) *Schedule {
	return &Schedule{Bot: bot}
}

func (s *Schedule) Start() {
	go s.Loop()
}

func (s *Schedule) Loop() {
	from := now() - tenMinutes
	var to int64
	for {
		to = now()
		mts := repository.GetSliceOfMTS(from, to)
		for _, e := range mts {
			s.sendReply(e)
		}
		from = to + 1
		time.Sleep(tenMinutesD)
	}

}

// private
func (s *Schedule) sendReply(messageToSend entity.MessageToSend) {
	forward := tgbotapi.NewForward(messageToSend.ChatID, messageToSend.ChatID, messageToSend.MessageID)
	if _, err := s.Bot.Send(forward); err != nil {
		log.Println(err)
		// TODO if err try to send it again in three minutes
		return
	}
	s.ContinueLoop(messageToSend)
}

func (s *Schedule) ContinueLoop(send entity.MessageToSend) {
	timeToSend := send.Data.NextIntervalTime + send.TimeToSend
	sendingNumOfData := send.SendingNumOfData + 1
	e := entity.ConstructorMTS(send.Data, sendingNumOfData, timeToSend)
	repository.SaveMessageToSend(e)
	repository.UpdateNextIntervalTime(send.Data)
}

func now() int64 {
	return time.Now().UnixNano()
}

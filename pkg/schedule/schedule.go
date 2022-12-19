package schedule

import (
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/entity"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"time"
)

type Schedule struct {
	Bot *tgbotapi.BotAPI
}

const (
	halfHour      = int64(30 * time.Minute)
	threeMinutes  = int64(3 * time.Minute)
	threeMinutesD = 3 * time.Minute
)

func NewSchedule(bot *tgbotapi.BotAPI) *Schedule {
	return &Schedule{Bot: bot}
}

func (s *Schedule) Start() {
	go s.Loop()
}

func (s *Schedule) Loop() {
	from := time.Now().UnixNano()
	var to int64
	for {
		to = from + halfHour
		mts := repository.GetSliceOfMTS(from, to)
		s.sendSchedule(mts, from, to)
		from = to + 1
	}

}

// private
func (s *Schedule) sendSchedule(mts []entity.MessageToSend, from, to int64) {
	index, l := 0, len(mts)
	for from < to {
		for index != l && mts[index].TimeToSend <= from+threeMinutes {
			s.sendReply(mts[index])
			index++
		}
		from += threeMinutes
		time.Sleep(threeMinutesD)
	}
	for index != l && mts[index].TimeToSend <= from+threeMinutes {
		s.sendReply(mts[index])
		index++
	}
}

func (s *Schedule) sendReply(messageToSend entity.MessageToSend) {
	forward := tgbotapi.NewForward(messageToSend.ChatID, messageToSend.ChatID, messageToSend.MessageID)
	if _, err := s.Bot.Send(forward); err != nil {
		//TODO
		return
	}
	repository.ContinueLoop(messageToSend)
}
func funcForSort(mts []entity.MessageToSend) func(i, j int) bool {
	return func(i, j int) bool {
		return mts[i].TimeToSend < mts[j].TimeToSend
	}
}

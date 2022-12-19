package schedule

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Schedule struct {
	Bot *tgbotapi.BotAPI
}

func NewSchedule(bot *tgbotapi.BotAPI) *Schedule {
	return &Schedule{Bot: bot}
}

func (s Schedule) Start() {
	go s.Loop()
}

func (s Schedule) Loop() {

}

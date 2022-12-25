package main

import (
	"fmt"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/cmd/getToken"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(getToken.TokenFromConsole())
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	telegram.Start(bot)
}

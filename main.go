package main

import (
	"fmt"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/dotEnv"
	"github.com/AbdullohAbdullayev/golang_integralli_takrorlash_bot.git/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	fmt.Println("\nstarting to get token")
	token := dotEnv.EnvMap["TOKEN"]
	fmt.Println("token is = ", token)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err)
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	telegram.Start(bot)
}

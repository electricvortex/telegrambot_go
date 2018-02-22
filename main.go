package main

import (
	"log"

	"strings"

	"math/rand"
	"time"

	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("515058850:AAGhdTqSCQ4NF9zkmziw__5d6oNi304uS04")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 2
	updates, err := bot.GetUpdatesChan(u)

	messages := []string{"Пиздаглазое мудоебище.", "Распиздяй колхозный.", "долбоебень", "блядюга", "ванючая хобитская сучка",
		"капитан потные яички", "ебучий хуебес", "Свиноблядь ебаная", "Говнарь", "Дичь обоссаная"}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		rand.Seed(time.Now().UTC().UnixNano())

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, messages[rand.Intn(len(messages)-1)])

		if strings.Contains(update.Message.Text, "Кини") {
			bot.Send(msg)
		} else if strings.Contains(update.Message.Text, "Попов") {
			bot.Send(msg)
		} else if strings.Contains(update.Message.Text, "кини") {
			bot.Send(msg)
		} else if strings.Contains(update.Message.Text, "попов") {
			bot.Send(msg)
		} else if strings.Contains(update.Message.Text, "Рауан") {
			bot.Send(msg)
		} else if strings.Contains(update.Message.Text, "рауан") {
			bot.Send(msg)
		}

	}
}

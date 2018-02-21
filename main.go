package main

import (
	"log"

	"math/rand"
	"strings"
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

	for update := range updates {
		if update.Message == nil {
			continue
		}

		messages := []string{"Петух!",
			"Пидераст!",
			"Говно!",
			"Чмо!",
			"Отстой!",
			"Лох!",
			"Тупица!",
			"Ванючка!"}

		if strings.Contains(update.Message.Text, "Кини") {
			rand.Seed(time.Now().UnixNano())

			choosen := messages[rand.Intn(len(messages)-1)]
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, choosen)
			bot.Send(msg)
		} else if strings.Contains(update.Message.Text, "Владимир") {
			rand.Seed(time.Now().UnixNano())

			choosen := messages[rand.Intn(len(messages)-1)]
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, choosen)
			bot.Send(msg)
		}

	}
}

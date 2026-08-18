package main

import (
	"log"
	"math/rand"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var userState = make(map[int64]Word) // userID → текущий Word

func main() {
	rand.Seed(time.Now().UnixNano())

	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.CallbackQuery != nil {
			handleCallback(bot, update.CallbackQuery)
			continue
		}

		if update.Message == nil {
			continue
		}

		chatID := update.Message.Chat.ID
		userID := update.Message.From.ID
		text := update.Message.Text

		if text == "/start" {
			showStartButton(bot, chatID)
			continue
		}

		if current, ok := userState[userID]; ok {
			HandleGameMessage(bot, chatID, userID, text, current)
			continue
		}
	}
}

func handleCallback(bot *tgbotapi.BotAPI, cb *tgbotapi.CallbackQuery) {
	userID := cb.From.ID
	chatID := cb.Message.Chat.ID

	if cb.Data == "start_game" {
		StartGame(bot, chatID, userID)
	}

	bot.Request(tgbotapi.NewCallback(cb.ID, ""))
}

func showStartButton(bot *tgbotapi.BotAPI, chatID int64) {
	button := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Начать игру", "start_game"),
		),
	)

	msg := tgbotapi.NewMessage(chatID, "Нажми кнопку, чтобы начать игру:")
	msg.ReplyMarkup = button

	bot.Send(msg)
}

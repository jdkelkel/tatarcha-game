package main

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isCorrectRussian(input string, w Word) bool {
	return strings.ToLower(strings.TrimSpace(input)) ==
		strings.ToLower(strings.TrimSpace(w.Russian))
}

func StartGame(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	current := RandomWord()

	bot.Send(tgbotapi.NewMessage(chatID,
		"Переведи слово:\nLatin: "+current.Latin+"\nCyril: "+current.Cyril))

	u := tgbotapi.NewUpdate(update.UpdateID + 1)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for upd := range updates {
		if upd.Message == nil {
			continue
		}

		userText := upd.Message.Text

		if isCorrectRussian(userText, current) {
			current = RandomWord()

			bot.Send(tgbotapi.NewMessage(chatID,
				"Верно! Новое слово:\nLatin: "+current.Latin+"\nCyril: "+current.Cyril))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID,
				"Неверно. Попробуй ещё раз."))
		}
	}
}

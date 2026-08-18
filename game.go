package main

import (
    "strings"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func isCorrectRussian(input string, w Word) bool {
    return strings.ToLower(strings.TrimSpace(input)) ==
        strings.ToLower(strings.TrimSpace(w.Russian))
}

func StartGame(bot *tgbotapi.BotAPI, chatID, userID int64) {
    current := RandomWord()
    userState[userID] = current

    msg := "Переведи слово:\nLatin: " + current.Latin + "\nCyril: " + current.Cyril
    bot.Send(tgbotapi.NewMessage(chatID, msg))
}

func HandleGameMessage(bot *tgbotapi.BotAPI, chatID, userID int64, text string, current Word) {
    if isCorrectRussian(text, current) {
        newWord := RandomWord()
        userState[userID] = newWord

        msg := "Верно! Новое слово:\nLatin: " + newWord.Latin + "\nCyril: " + newWord.Cyril
        bot.Send(tgbotapi.NewMessage(chatID, msg))
    } else {
        bot.Send(tgbotapi.NewMessage(chatID, "Неверно. Попробуй ещё раз."))
    }
}

package main

import (
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartGame(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

    chatID := update.Message.Chat.ID

    current := RandomWord()

    bot.Send(tgbotapi.NewMessage(chatID,
        "Переведи слово:\nLatin: "+current.Latin+"\nCyril: "+current.Cyril))
}

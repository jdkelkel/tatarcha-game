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
        if update.Message == nil {
            continue
        }

        chatID := update.Message.Chat.ID
        userID := update.Message.From.ID
        text := update.Message.Text


        if text == "/startgame" {
            StartGame(bot, chatID, userID)
            continue
        }


        if current, ok := userState[userID]; ok {
            HandleGameMessage(bot, chatID, userID, text, current)
            continue
        }
    }
}

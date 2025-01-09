package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/go-redis/redis/v8"
    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
    bot         *tgbotapi.BotAPI
    redisClient  *redis.Client
    ctx         = context.Background()
)

func init() {
    var err error
    bot, err = tgbotapi.NewBotAPI("7717459803:AAEsKjO39AIeoVLtxZs9Fv7x8yrKlvpOn5A")
    if err != nil {
        log.Panic(err)
    }

    redisClient = redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_ADDR"), 
        Password: "", 
        DB:       0,  
    })
}

func main() {
    log.Printf("Authorized on account %s", bot.Self.UserName)

   
    updates := bot.ListenForWebhook("/" + "7717459803:AAEsKjO39AIeoVLtxZs9Fv7x8yrKlvpOn5A")

    go func() {
        log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
    }()

    for update := range updates {
        if update.Message == nil { 
            continue
        }

        chatID := update.Message.Chat.ID

      
        if isUserRegistered(chatID) {
            msg := tgbotapi.NewMessage(chatID, "Вы зарегистрированы! Ваше сообщение: "+update.Message.Text)
            bot.Send(msg)
        } else {
            msg := tgbotapi.NewMessage(chatID, "Вы не зарегистрированы. Пожалуйста, зарегистрируйтесь.")
            bot.Send(msg)
        }
    }
}


func isUserRegistered(chatID int64) bool {
    exists, err := redisClient.Exists(ctx, fmt.Sprintf("user:%d", chatID)).Result()
    if err != nil {
        log.Println("Ошибка проверки пользователя:", err)
        return false
    }
    return exists > 0
}